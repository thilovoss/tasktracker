<script setup lang="ts">
import TaskItem from "./TaskItem.vue"
import AddTask from "./AddTask.vue"
import { useTaskStore } from '../stores/tasks'
import { computed } from "@vue/reactivity";
import { onMounted } from "vue";

const store = useTaskStore()

const getTasks = computed(() => {
    return store.getTasks
})

const tasks = computed(() => {
    return store.tasks
})

onMounted(() => {
    store.fetchTasks()
})
</script>

<template>
    <div v-for='gettersTask in getTasks'>
        <TaskItem :task="gettersTask"></TaskItem>
    </div>
</template>