<script setup lang="ts">
const store = useTaskStore()

async function deleteTask(id: String) {
    await store.deleteTask(id)
    store.fetchTasks()
}

function calculateDuration(start: String, end: String) {
    const startDate = new Date(start)
    const endDate = new Date(end)
    const differenceMilliseconds = endDate - startDate
    var difference: String
    if (differenceMilliseconds < 3.6e6) {
        difference = differenceMilliseconds / 60000 + "m"
    } else if (differenceMilliseconds < 8.64e7) {
        difference = differenceMilliseconds / 3.6e6 + "h"
    } else {
        difference = differenceMilliseconds / 8.64e7 + "d"
    }
    // const minutes = (endDate - startDate) / 60000

    return difference
}

function convertDate(dateString: String) {
    const date: Date = new Date(dateString)
    const dateOptions = {
        dateStyle: "short",
    }
    const timeOptions = {
        timeStyle: "short",
    }
    return date.toLocaleDateString([], dateOptions) + " " + date.toLocaleTimeString([], timeOptions)
}
</script>

<script lang="ts">
import { useTaskStore, type Task } from '../stores/tasks'
export default {
    props: { task: null }
}
</script>

<template>
    <div class="grid grid-cols-5  border-solid border-2 m-2 rounded-md p-2">
        <div class="col-span-2">
            <h2>{{ task.name }}</h2>
            <p>{{ task.description }}</p>
            
        </div>
        <div class="col-span-1">
            <div>start: {{ convertDate(task.start) }}</div>
            <div>end: {{ convertDate(task.end) }}</div>
            <div>duration: {{ calculateDuration(task.start, task.end) }}</div>
        </div>
        <button v-on:click="deleteTask(task.id)" class="col-span-2 my-auto justify-self-end">delete</button>
    </div>
</template>