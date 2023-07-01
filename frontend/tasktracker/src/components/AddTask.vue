
<script setup lang="ts">
import { useCategoryStore, type Category } from '@/stores/categories';
import { useTaskStore, type Task } from '@/stores/tasks'
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'
import { computed, onMounted, ref } from 'vue';

const taskStore = useTaskStore()
const descriptionTextarea = ref<HTMLElement>()

const categoryStore = useCategoryStore()
const getCategories = computed(() => {
    return categoryStore.getCategories
})

onMounted(() => {
    categoryStore.fetchCategories()
})

function addTask(data: any) {
    console.log(data)
    const task: Task = {
        name: data.name,
        description: data.description,
        start: data.date[0],
        end: data.date[1],
        categoryId: "4e256b3d-8162-4676-8d43-cd888e806c5d"
    }
    taskStore.createTask(task).then(() => { taskStore.fetchTasks() })
}

// TODO: resize funktioniert noch nicht richtig
function resize() {
    let element = descriptionTextarea.value;
    element!.style.height = "auto";
    element!.style.height = element!.scrollHeight + "px";
}
</script>

<script lang="ts">
export default {
    components: { VueDatePicker },
    data() {
        return {
            date: null,
            name: null,
            description: null
        };
    }
}

</script>

<template>
    <div class="grid grid-cols-5 p-2">
        <input v-model="name" type="text" placeholder="title" class="col-span-1 mr-2 my-0 h-9 text-primary" />
        <textarea :maxlength="150" v-model="description" type="textarea" @input="resize()" ref="descriptionTextarea"
            placeholder="descripton" class="col-span-2 mr-2 overflow-hidden resize-none h-9 pt-1 text-primary"></textarea>
        <select>
            <option v-for="category in getCategories" :key="category.id" :value="category.id">
                {{ category.title }}</option>
        </select>
        <VueDatePicker v-model="date" range class="max-w-md col-span-2"></VueDatePicker>
        <button v-on:click="addTask($data)" class="justify-self-end col-span-5 text-primary my-auto">add</button>
    </div>
</template>