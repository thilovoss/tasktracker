<script lang="ts">
import type { Task } from '../stores/tasks'
export default {
    props: {task: Object as unknown as Task}
}
</script>

<template>
    <div>
        <h2>{{task.name}}</h2>
        <p>{{task.description}}</p>
        <div>start: {{task.start}}</div>
        <div>end: {{task.end}}</div>
    </div>
</template>