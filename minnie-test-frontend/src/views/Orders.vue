<template>
  <div class="container">
    <SearchBar @update-search-text="updateSearchText"/>
    <DateFilter @update-search-dates="updateSearchDates"/>

    <p :style="{'margin-top': '30px'}">Total amount <span :style="{'font-weight': 'bold'}">${{totalAmount.toFixed(2)}}</span></p>
 
    <div :style="{'margin-top': '20px'}">

    <Table
      :rowsPerPage="rowsPerPage"
      :currentPageNum="currentPageNum"
      :totalPages="totalPages"
      :rows = "orders"
      @go-to-previos-page="goToPreviosPage"
      @go-to-next-page="goToNextPage"
      @jump-to-page="jumpToPage"
      @rows-per-page-changed="changeRowsPerPage"
      @sort-rows-by-date="sortRows"
    />
    </div>
 </div>
</template>

<script>
import moment from 'moment'
import SearchBar from '../components/SearchBar';
import DateFilter from '../components/DateFilter';
import Table from '../components/Table';
import utils from '../utils';
export default {
  name: 'OrdersComponent',
  components:{
    SearchBar,
    DateFilter,
    Table
  },
  created() {
    this.searchText = 'all'
    this.makeRequestAndUpdateTable()
  },
	data() {
		return {
      searchText: null,
      startDate: null,
      endDate: null,
      totalAmount: 198,

      rowsPerPage: 5,
      currentPageNum: 1,
      totalPages: 5,

      sortedType: null,

      getCount: false,

      orders: [],
		}
	},
  watch: {
    searchText: async function(newVal, oldVal) {
      if (newVal !== oldVal){
        console.log('searchText CHANGE TRIGGERED')
        this.getCount = true
        this.jumpToPage(1)
        this.makeRequestAndUpdateTable()
        this.getCount = false
      }
    },
    startDate: async function(newVal, oldVal) {
      if (newVal !== oldVal){
        this.getCount = true
        this.jumpToPage(1)
        await this.makeRequestAndUpdateTable()
        this.getCount = false
      }
    },
    endDate: async function(newVal, oldVal) {
      if (newVal !== oldVal){
        this.getCount = true
        this.jumpToPage(1)
        await this.makeRequestAndUpdateTable()
        this.getCount = false
      }
    },
    rowsPerPage: async function(newVal, oldVal) {
      if (newVal !== oldVal){
        this.getCount = true
        this.jumpToPage(1)
        await this.makeRequestAndUpdateTable()
        this.getCount = false
      }
    },
    currentPageNum: async function(newVal, oldVal) {
      if (newVal !== oldVal){
        this.getCount = true
        await this.makeRequestAndUpdateTable()
        this.getCount = false
      }
    }
  },
	methods: {
    sortRows() {
      if (this.sortedType == null || this.sortedType == 'ascend' ){
        this.orders.sort((a, b) => {
          let time1 = moment(a.orderDate, 'YYYY-MM-DDThh:mm:ss').parseZone("Australia/Melbourne").valueOf();
          let time2 = moment(b.orderDate, 'YYYY-MM-DDThh:mm:ss').parseZone("Australia/Melbourne").valueOf();
          return time1 - time2
        })
        this.sortedType = 'descend'
      } else {
        this.orders.sort((a, b) => {
          let time1 = moment(a.orderDate, 'YYYY-MM-DDThh:mm:ss').parseZone("Australia/Melbourne").valueOf();
          let time2 = moment(b.orderDate, 'YYYY-MM-DDThh:mm:ss').parseZone("Australia/Melbourne").valueOf();
          return time2 - time1
        })
        this.sortedType = 'ascend'
      }
    },
    goToPreviosPage() {
      if (this.currentPageNum >= 2) {
        this.currentPageNum -= 1
        this.makeRequestAndUpdateTable()
      }
    },
    goToNextPage() {
      if (this.currentPageNum < this.totalPages){
        this.currentPageNum += 1
        this.makeRequestAndUpdateTable()
      } else {
        this.currentPageNum = this.totalPages 
      }
    },
    jumpToPage(pageValue) {
      console.log(pageValue)
      if (pageValue <= this.totalPages){
        this.currentPageNum = Number(pageValue)
        this.makeRequestAndUpdateTable()
      } else {
        this.currentPageNum = this.totalPages
        this.makeRequestAndUpdateTable()
      }
    },
    changeRowsPerPage(newVal) {
      console.log({newVal})
    },
    async makeRequestAndUpdateTable() {
      try {       
        if (!this.searchText){
          this.searchText = ''
        }
        if (!this.startDate){
          this.startDate = '1022-01-01'
        }
        if (!this.endDate){
          this.endDate = '3022-01-01'
        }
        if (!this.sortedType){
          this.sortedType = 'ascend'
        }
       
        let response = await fetch(
          utils.BASE_URL + 
          "/orders/p/" + this.searchText + 
          "/o/" + this.searchText +
          "/s/" + this.startDate +
          "/e/" + this.endDate +
          "/s/" + this.sortedType +
          "/p/" + this.currentPageNum +
          "/c/" + this.getCount
        )
        response = await response.json()
        if (response.status === 'success'){
          let newOrders = []
          this.totalAmount = 0
          newOrders = response.data.map((dataObj, index) => {
            let { company_name: customerCompany, customer_name: customerName, delivered_amount: deliveredAmount, order_name: orderName, order_time: orderDate, total_amount: total } = dataObj
            this.totalAmount += total
            return { id: String(index), customerCompany, customerName, deliveredAmount, orderName, orderDate, total }
          })
          this.orders = newOrders
          if (response.totalRows) {
            this.totalPages = Math.ceil(response.totalRows / 5)
          } 
        } else {
          this.orders = []
          this.totalPages = 1
        }
      } catch (err) {
        this.orders = []
        this.totalPages = 1
        console.log(err)
      }      
    },
		async updateSearchText(searchText) {
      this.searchText = searchText
		},

		updateSearchDates(startDate, endDate) {
      this.startDate = startDate
      this.endDate = endDate
		},
  },
};
</script>

<style>
.container {
  margin-top: 10px;
  width: '80%';
  border-radius: 2px;
  border-color: #eee;
  border-style: solid;
  border-width: 1px;
  padding: 20px;
  padding-bottom: 80px;
}
.table-header-container {
  margin-top: 20px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
.table-column {
  flex: 1;
  text-align: left;
}
</style>
