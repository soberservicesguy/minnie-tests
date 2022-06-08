<template>
	<div class='form'>
		<div class='search-label-container'>
			<i class='fa fa-search'></i><label class='search-label'>Search</label>
		</div>
		<div class='search-bar-container'>
			<input class='search-bar' type='text' v-model='searchText' @change='onChange' name='text' placeholder='please enter something to search' />
		</div>
	</div>

</template>

<script>
export default {
	name: 'SearchBar',
	data() {
		return {
			searchText: '',
		}
	},
	methods: {
		onChange(e) {
			e.preventDefault()
			if (!this.searchText) {
				this.$emit('update-search-text', "all")
				this.searchText = ""
			} else {
				this.$emit('update-search-text', this.searchText)
			}
		},
	},
};
</script>

<style  scoped>
.search-label {
	font-size: 1.2rem;
	margin-left: 10px;
}
.search-bar {
	width: 100%;
	border-width: 0;
	border-right-width: 1px;
	border-color: #eee;

}
.form {
	width: 100%;
	display: flex;
	flex-direction: row;
}
.search-bar-container {
	margin-left: 40px;
	flex: 10;
	height: 20px;
	padding: 2px;
	padding-left: 5px;
	border-radius: 5px;
	border-width: 1px;
	border-width: 1px;
	border-color: #eee;
	border-style: solid;
	margin-right: 50px;
}
.search-label-container {
	flex: 1;
}
</style>
