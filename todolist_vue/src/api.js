import axios from 'axios'
import adapter from "axios/lib/adapters/http";

axios.defaults.adapter = adapter;

export class API {
    constructor() {
        this.url = "http://localhost:8080/api/v1"
    }
    
    async getTodoList() {
        return axios.get(this.url+'/getTodoList').then(r => r.data)
    }

    async addTodo(todo) {
        return axios.post(this.url+'/addTodo', {task_description: todo}).then(r => r.data)
    }
}

export default new API();