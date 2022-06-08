import { mount } from '@vue/test-utils'
import DateFilter from '../../../../src/components/DateFilter'

let values = [
  {startDate: '2020-10-10', endDate: '2020-10-11'},
  {startDate: '2020-10-11', endDate: '2020-10-12'},
  {startDate: '2020-01-10', endDate: '2020-01-30'},
  {startDate: '2020-01-10', endDate: '2020-01-20'},
  {startDate: '2020-01-10', endDate: '2020-02-11'},
  {startDate: '2020-02-10', endDate: '2020-03-11'},
  {startDate: '2020-03-10', endDate: '2020-04-11'},
]

const getStartAndEndDate = () => {
  const wrapper = mount(DateFilter)
  const startDateInput = wrapper.find("input[placeholder='Start date']")
  const endDateInput = wrapper.find("input[placeholder='End date']")
  return { wrapper, startDateInput, endDateInput }
}

const enterStartAndEndDate = async (startDateInput, endDateInput, startDate, endDate) => {
  await startDateInput.setValue(startDate)
  await endDateInput.setValue(endDate)
}

const testStartAndEndDateInStateAndDOM = async (startDate, endDate) => {
  const { wrapper, startDateInput, endDateInput } = getStartAndEndDate()
  await enterStartAndEndDate(startDateInput, endDateInput, startDate, endDate)

  // DOM value
  expect(startDateInput.element.value).toBe(startDate)
  expect(endDateInput.element.value).toBe(endDate)

  // state
  expect(wrapper.vm.$data.startDate).toBe(startDate)
  expect(wrapper.vm.$data.endDate).toBe(endDate)
}

const testStartAndEndDateInEvent = async (startDate, endDate) => {
  const { wrapper, startDateInput, endDateInput } = getStartAndEndDate()
  await enterStartAndEndDate(startDateInput, endDateInput, startDate, endDate)

  const datesEntryEvent = wrapper.emitted('update-search-dates')
  
  // total event fires
  expect(datesEntryEvent).toHaveLength(2)

  // event payload
  expect(datesEntryEvent[0]).toEqual([startDate, ''])
  expect(datesEntryEvent[1]).toEqual([startDate, endDate])
} 


describe('DateFilter.vue updates startDate and endDate both on state and DOM and passes them in event', () => {
  it('updates startDate and endDate in state and DOM', async () => {
    let testPromises = values.map(async (valueObj) => testStartAndEndDateInStateAndDOM(valueObj.startDate, valueObj.endDate))
    await Promise.all(testPromises)    
  })

  it('fires update-search-dates event with new startDate and endDate', async () => {
    await testStartAndEndDateInEvent('2020-10-10', '2020-10-11')
    await testStartAndEndDateInEvent('2020-10-11', '2020-10-12')
    await testStartAndEndDateInEvent('2020-10-12', '2020-10-14')
  })

})
