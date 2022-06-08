<template>
	<div class='form'>
		<label class='search-label'>Created date</label><br/>
		<div class='date-container'>
			<div class='start-date-container'>
				<img src='../assets/calendar.png' alt='drop down icon' width=15 height=15 :style="{'margin-left': '5px'}"/>
				<input 
					class='date-input'
					type='text'
					v-model='startDate'
					@change='onChange'
					name='startDate'
					placeholder='Start date' 
					onfocus="(this.type='date')"
				/>
				<img v-show='showDropDownIcon' src='../assets/drop-down-arrow.png' alt='drop down icon' width=20 height=20/>
			</div>
			<div class='end-date-container'>
				<input 
					class='date-input'
					type='text'
					v-model='endDate'
					@change='onChange'
					name='endDate'
					placeholder='End date' 
					onfocus="(this.type='date')"
				/>
			</div>
		</div>

	</div>

</template>

<script>
export default {
	name: 'DateFilter',
	data() {
		return {
			startDate: '',
			endDate: '',
		}
	},
	computed: {
		showDropDownIcon() {
			if (this.startDate == ''){
				return true
			} else {
				return false
			}
		}
	},
	methods: {
		onChange(e) {
			e.preventDefault()
			this.$emit('update-search-dates', this.startDate, this.endDate)
		},
	},
};
</script>

<style scoped>
.form {
	padding-top: 20px;
}
icon {
    width: 2px;
    height: 2px;
}
input:focus{
    outline: none;
}
.date-container {
	margin-top: 10px;
	height: 30px;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    width: 320px;
	border-width: 1px;
	border-color: #eee;
	border-style: solid;
	border-radius: 4px;
}
.start-date-container {
    display: flex;
    align-items: center;
    justify-content: flex-start;
}
.end-date-container {
	padding-left: 10px;
}
.search-label {
	font-size: 20px;
}
.date-input {
	height: 20px;
	width: 120px;
	margin-left: 10px;
    border-width: 0;
}
</style>
