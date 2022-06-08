import { mount } from '@vue/test-utils'
import TableRow from '../../../../src/components/TableRow'
import moment from 'moment'
let values = [
    {
        orderName: "abc",
        customerCompany: "cisco",
        customerName: "jane",
        orderDate: "2020-10-10T00:34:45",
        deliveredAmount: 123,
        total: 124,
    },
    {
        orderName: "abc",
        customerCompany: "cisco",
        customerName: "jane",
        orderDate: "2020-10-10T00:34:45",
        deliveredAmount: 123,
        total: 124,
    },
    {
        orderName: "abc",
        customerCompany: "cisco",
        customerName: "jane",
        orderDate: "2020-10-10T00:34:45",
        deliveredAmount: 123,
        total: 124,
    },
]

const getTableRowContents = () => {
  const wrapper = mount(TableRow)
  const orderNameCell = wrapper.find("#col-1")
  const companyNameCell = wrapper.find("#col-2")
  const customerNameCell = wrapper.find("#col-3")
  const orderDateCell = wrapper.find("#col-4")
  const deliveredAmountCell = wrapper.find("#col-5")
  const totalAmountCell = wrapper.find("#col-6")
  return { wrapper, orderNameCell, companyNameCell, customerNameCell, orderDateCell, deliveredAmountCell, totalAmountCell }
}

const setProps = async (wrapper, propsObj) => {
    await wrapper.setProps(propsObj)
}

const testTableRowPropsInDOM = async (propsObj) => {
    const { wrapper, orderNameCell, companyNameCell, customerNameCell, orderDateCell, deliveredAmountCell, totalAmountCell } = getTableRowContents()
    await setProps(wrapper, propsObj)

    // DOM value
    expect(orderNameCell.text()).toBe(propsObj.orderName)
    expect(companyNameCell.text()).toBe(propsObj.customerCompany)
    expect(customerNameCell.text()).toBe(propsObj.customerName)
    expect(orderDateCell.text()).toBe(moment(propsObj.orderDate, 'YYYY-MM-DDThh:mm:ss').format('MMM Do, h:mm A'))
    expect(deliveredAmountCell.text()).toBe("$" + propsObj.deliveredAmount)
    expect(totalAmountCell.text()).toBe("$" + propsObj.total)
}

describe('TableRow.vue updates props supplied into DOM and shows dates, currency properly', () => {
  it('updates props into DOM', async () => {
    let testPromises = values.map(async (propObj) => testTableRowPropsInDOM(propObj))
    await Promise.all(testPromises)    
  })
})