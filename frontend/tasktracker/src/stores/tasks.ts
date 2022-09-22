import { defineStore } from 'pinia'
import axios from "axios"

export interface Task {
    name: string;
    description: string;
    start: string;
    end: string;
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
        }
    },
})