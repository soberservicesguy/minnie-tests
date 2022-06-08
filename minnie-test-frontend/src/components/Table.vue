<template>
    <div class='table-row-container'>
      <div class='table-column' :style="{'flex': 2,}"><p>Order name</p></div>
      <div class='table-column' :style="{'flex': 1.5}"><p>Customer Company</p></div>
      <div class='table-column' :style="{'flex': 1.5}"><p>Customer name</p></div>
      <button @click="this.$emit('sort-rows-by-date')" class='table-column' :style="{'background':'none', 'border': 'none', 'display': 'flex', 'align-items': 'center', 'justify-content': 'flex-start', 'flex': '1.5'}">
        <p>
            Order date
        </p>
            <img
                src='../assets/sort-icon.png'
                alt='previour page button'
                width='10'
                height='10'
                :style="{'margin-left': '5px'}"
            />
        </button>
      <div class='table-column'><p>Delivered Amount</p></div>
      <div class='table-column' :style="{'text-align': 'left'}"><p>Total Amount</p></div>
    </div>

    <div :key='rowData.id' v-for='rowData in rows'>
      <TableRow
        :orderName='rowData.orderName'
        :customerCompany='rowData.customerCompany'
        :customerName='rowData.customerName'
        :orderDate='rowData.orderDate'
        :deliveredAmount='rowData.deliveredAmount'
        :total='rowData.total'
      />
    </div>

    <div class='horizontal-container'>
        <p id="total-pages">
            Total {{totalPages}}
        </p>

        <select class='page-count-dropdowns' v-model="rowsNum" @change="this.$emit('rows-per-page-changed', this.rowsNum)" name='rowCounts'>
            <option value='5'>5/page</option>
            <option value='25'>25/page</option>
        </select>

        <div class='horizontal-container' :style="{'margin-left': '20px'}">
            <button id="go-previos-page" class="go-to-page-button" @click="this.$emit('go-to-previos-page')">
                <img src='../assets/left-arrow.png' alt='previous page button' width='10' height='10' :style="{'margin-left': '5px'}"/>
            </button>
            <p class='current-page'>
                {{currentPageNum}}
            </p>
            <button id="go-next-page" class="go-to-page-button" @click="this.$emit('go-to-next-page')">
                <img src='../assets/right-arrow.png' alt='next page button' width='10' height='10' :style="{'margin-left': '5px'}"/>
            </button>
        </div>

        <div class='horizontal-container' :style="{'margin-left': '20px'}">
            <p>Go to</p>
            <input
                class='go-to-page-input'
                type='text'
                v-model="pageNum"
                @change="this.$emit('jump-to-page', pageNum)" 
                name='text'
                placeholder='please enter something to search'
            />
        </div>

    </div>

</template>

<script>
import TableRow from './TableRow.vue'
export default {
    name: 'TableComponent',
    props: {
        rows:Array,
        currentPageNum: {
            type: Number,
            default: 1
        },
        rowsPerPage: {
            type: Number,
            default: 5,
        },
        totalPages: {
            type: Number,
            default: 5,
        }
    },
    watch: { 
        currentPageNum: function(newVal, oldVal) {
            if (newVal !== oldVal){
                this.pageNum = newVal
            }
        },
        rowsPerPage: function(newVal, oldVal) {
            if (newVal !== oldVal){
                this.rowsNum = newVal
            }
        },
    },
    components:{
        TableRow,
    },
	data() {
		return {
            pageNum: this.currentPageNum,
            rowsNum: this.rowsPerPage,
        }
	},
	methods: {
        async searchOrders(searchText) {
            this.searchText = searchText
		},
		async filterOrders(startDate, endDate) {
            this.startDate = startDate
            this.endDate = endDate
        },

    },
};
</script>

<style scoped>
.go-to-page-button {
    background: none;
    border: none;
}
.table-row-container {
  margin-top: 20px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
.table-column {
  flex: 1;
  text-align: left;
  color: grey;
}
.page-count-dropdowns {
    text-align: center;
    margin-left: 20px;
    background-color: white;
    border: 1px solid #eee;
    border-radius: 4px;
    padding-left: 5px;
    padding-top: 2px;
    padding-bottom: 2px;
}
.horizontal-container {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    
}
.current-page {
    margin-left: 14px;
    margin-right: 10px;
    border: 1px solid #EEE;
    padding: 5px;
    padding-left: 10px;
    padding-right: 10px;
    border-radius: 5px;
}
.go-to-page-input {
    width: 50px;
    margin-left: 5px;
    border: 1px solid #eee;
    border-radius: 5px;
    text-align: center;
    padding-top: 5.5px;
    padding-bottom: 5.5px;
}
</style>
