
<script setup lang="ts">
import { useTaskStore, type Task } from '@/stores/tasks'
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'
import { onMounted, ref } from 'vue';

const store = useTaskStore()

function addTask(data: any) {
    console.log(data)
    const task: Task = {
        name: data.name,
        description: data.description,
        start: data.date[0].toISOString(),
        end: data.date[1].toISOString(),
        categoryId: "4e256b3d-8162-4676-8d43-cd888e806c5d"
    }
    store.createTask(task)
}
</script>

<script lang="ts">
export default {
    components: { VueDatePicker },
    data() {
        return {
            date: null,
            name: null,
            description: null,
        };
    }
}



</script>

<template>
    <input v-model="name" type="text" placeholder="title" />
    <input v-model="description" type="text" placeholder="descripton" />
    <VueDatePicker v-model="date" range></VueDatePicker>
    <button v-on:click="addTask($data)">add</button>
</template>