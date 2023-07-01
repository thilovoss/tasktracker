import { defineStore } from 'pinia'
import axios from "axios"

export interface Task {
    id?: String;
    name: string;
    description: string;
    start: Date;
    end: Date;
    categoryId?: string;
}


export const useTaskStore = defineStore("tasks", {
    state: () => ({
        tasks: [] as Task[]
    }),
    getters: {
        getTasks(state) {
            return state.tasks
        }
    },
    actions: {
        async fetchTasks() {
            try {
                // const data = await axios.get(import.meta.env.TASK_ENDPOINT!)
                const data = await axios.get("http://localhost:8080/api/task")
                console.log(data.status)
                this.tasks = data.data
            } catch (error) {
                console.log(this.getTasks)
                console.log(error)
            }
        },
        async createTask(task: Task) {
            try {
                await axios.post("http://localhost:8080/api/task", task)
            } catch (error) {
                console.log(error)
            }
        },
        async deleteTask(id: String) {
            try {
                await axios.delete("http://localhost:8080/api/task/"+id)
            } catch (error) {
                console.log(error)
            }
        }
    },
})