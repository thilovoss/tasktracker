<script setup lang="ts">
    import TaskItem from "./TaskItem.vue"
</script>

<template>
    <div v-for="n in 10">
        <TaskItem></TaskItem>
    </div>  
</template>