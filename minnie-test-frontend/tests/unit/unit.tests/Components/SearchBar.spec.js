import { mount } from '@vue/test-utils'
import SearchBar from '../../../../src/components/SearchBar'

let values = [
    'ani',
    'mas',
    'sti',
    'gel',
    'coke',
]
  
const getSearchBar = () => {
    const wrapper = mount(SearchBar)
    const searchInput = wrapper.find("input[placeholder='please enter something to search']")
    return { wrapper, searchInput }
}

const enterSearch = async (searchInput, value) => {
    await searchInput.setValue(value)
}
  

const testSearchInputInStateAndDOM = async (value) => {
    const { wrapper, searchInput } = getSearchBar()
    await enterSearch(searchInput, value)
  
    // DOM value
    expect(searchInput.element.value).toBe(value)
  
    // state
    expect(wrapper.vm.$data.searchText).toBe(value)
}

const testSearchTextInEvent = async (value) => {
    const { wrapper, searchInput } = getSearchBar()
    await enterSearch(searchInput, value)
  
    const searchEntryEvent = wrapper.emitted('update-search-text')
  
    // event payload
    expect(searchEntryEvent[0]).toEqual([value])
  } 

describe('SearchBar.vue updates searchText both on state and DOM and passes them in event', () => {
    it('updates startDate and endDate in state and DOM', async () => {
      let testPromises = values.map(async (value) => testSearchInputInStateAndDOM(value))
      await Promise.all(testPromises)    
    })

    it('fires update-search-dates event with searchText', async () => {
        await testSearchTextInEvent('ani')
        await testSearchTextInEvent('mas')
        await testSearchTextInEvent('gel')
    })

    it('fires update-search-dates event with "all" if empty searchText', async () => {
        const { wrapper, searchInput } = getSearchBar()
        await enterSearch(searchInput, '')
      
        const searchEntryEvent = wrapper.emitted('update-search-text')
      
        // event payload
        expect(searchEntryEvent[0]).toEqual(['all'])    
    })

})