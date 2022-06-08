import { mount } from '@vue/test-utils'
import TableComponent from '../../../../src/components/Table'

const getTableContent = () => {
    const wrapper = mount(TableComponent)
    const goToPageInput = wrapper.find(".go-to-page-input")
    const goToPreviousPageButton = wrapper.find("#go-previos-page")
    const goToNextPageButton = wrapper.find("#go-next-page")
    const totalPages = wrapper.find("#total-pages")
    const currentPage = wrapper.find(".current-page")
    return { wrapper, goToPageInput, goToPreviousPageButton, goToNextPageButton, totalPages, currentPage }
}

const enterPageInput = async (goToPageInput, value) => {
    await goToPageInput.setValue(value)
}

const clickButton = async (button) => {
    await button.trigger('click')
}


const testPageNumInputForEvent = async (goToPageNum) => {
    const { wrapper, goToPageInput } = getTableContent()
    await enterPageInput(goToPageInput, goToPageNum)    

    const pageNumEvent = wrapper.emitted('jump-to-page')
    expect(pageNumEvent[0]).toEqual([String(goToPageNum)])
    expect(wrapper.vm.$data.pageNum).toBe(String(goToPageNum))
}

const testClickingGoToPrevPageButtonForEvent = async () => {
    const { wrapper, goToPreviousPageButton } = getTableContent()
    clickButton(goToPreviousPageButton)
    const clickPrevButtonEvent = wrapper.emitted('go-to-previos-page')
    expect(clickPrevButtonEvent[0]).toEqual([])
}

const testClickingGoToNextPageButtonForEvent = async () => {
    const { wrapper, goToNextPageButton } = getTableContent()
    clickButton(goToNextPageButton)
    const clickPrevButtonEvent = wrapper.emitted('go-to-next-page')
    expect(clickPrevButtonEvent[0]).toEqual([])
}


describe('Table.vue emits event when enter current page input', () => {
    it('emits event when entered in input', async () => {
        await testPageNumInputForEvent(5)
    })
})

describe('Table.vue emits events when click either go to prev or next page butotn', () => {
    it('emits event when clicked', async () => {
        await testClickingGoToPrevPageButtonForEvent()
        await testClickingGoToNextPageButtonForEvent()
    })
})