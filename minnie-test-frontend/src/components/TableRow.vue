<template>
    <div class='table-row-container'>
      <div class='table-column' :style="{'flex': 2}"><p id="col-1">{{orderName}}</p></div>
      <div class='table-column' :style="{'flex': 1.5}"><p id="col-2">{{customerCompany}}</p></div>
      <div class='table-column' :style="{'flex': 1.5}"><p id="col-3">{{customerName}}</p></div>
      <div class='table-column' :style="{'flex': 1.5}"><p id="col-4">{{getProperDate(orderDate)}}</p></div>
      <div class='table-column'><p id="col-5">{{deliveredAmount ? `$${deliveredAmount}`: ''}}</p></div>
      <div class='table-column' :style="{'text-align': 'left'}"><p id="col-6">{{total ? `$${total}`: ''}}</p></div>
    </div>
</template>

<script>
import moment from 'moment'
export default {
    name: 'TableRow',
    props: {
        orderName: {
            type: String,
            default: ''
        },
        customerCompany: {
            type: String,
            default: ''
        },
         customerName: {
            type: String,
            default: ''
        },
        orderDate: {
            type: String,
            default: ''
        },
        deliveredAmount: {
            type: Number,
            default: 0
        },
        total: {
            type: Number,
            default: 0
        },
    },
    components:{
    },
	data() {
		return {
		}
	},
	methods: {
        getProperDate(orderDate){
            return moment(orderDate, 'YYYY-MM-DDThh:mm:ss').format('MMM Do, h:mm A')
        },
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
.table-row-container {
  margin-top: 20px;
  display: flex;
  flex-diretion: row;
  align-items: center;
  justify-content: center;
  border-bottom: 1px solid #EEE;
}
.table-column {
  flex: 1;
  text-align: left;
}
</style>
