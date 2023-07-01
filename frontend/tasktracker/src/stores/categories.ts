import axios from "axios";
import { defineStore } from "pinia";

export interface Category {
    id?: String;
    title: String;
}

export const useCategoryStore = defineStore("categories", {
    state: () => ({
        categories: [] as Category[]
    }),
    getters: {
        getCategories(state) {
            return state.categories
        }
    },
    actions: {
        async fetchCategories() {
            try {
                const data = await axios.get("http://localhost:8080/api/category")
                console.log(data.status)
                this.categories = data.data
            } catch (error) {
                console.log(this.categories)
                console.log(error)
            }
        }
    }
})