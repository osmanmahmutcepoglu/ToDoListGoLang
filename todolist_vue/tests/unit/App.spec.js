import {shallowMount} from "@vue/test-utils";
import flushPromises from "flush-promises";
import App from "@/App";

jest.mock("@/api", () => ({
    getTodoList: () => Promise.resolve({data: []}),
    addTodo: (todo) => Promise.resolve({data: {id: 0, description: 'TEST 1'}})
}));

describe('App.vue', () => {
    it('Add', async () => {
        const inputToAdd = "TEST 1"

        const wrapper = shallowMount(App)
        await flushPromises();

        const input = wrapper.find('input')
        await input.setValue(inputToAdd)

        const button = wrapper.find('#addTodo')
        await button.trigger('click')
        await flushPromises();

        expect(wrapper.vm.todo).toBe('')
        expect(wrapper.vm.todoList[0].id).toEqual(0)
        expect(wrapper.vm.todoList[0].description).toBe(inputToAdd)
        expect(wrapper.find('p').text()).toBe('1. ' + inputToAdd)
    })
})