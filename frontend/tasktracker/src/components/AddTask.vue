
<script setup lang="ts">
import { useTaskStore, type Task } from '@/stores/tasks'
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'
import { ref } from 'vue';

const store = useTaskStore()
const descriptionTextarea = ref<HTMLElement>()

function addTask(data: any) {
    console.log(data)
    const task: Task = {
        name: data.name,
        description: data.description,
        start: data.date[0],
        end: data.date[1],
        categoryId: "4e256b3d-8162-4676-8d43-cd888e806c5d"
    }
    store.createTask(task).then(() => { store.fetchTasks() })
}

// TODO: resize funktioniert noch nicht richtig
function resize() {
    let element = descriptionTextarea.value;
    element!.style.height = "2.25rem";
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
            description: null,
        };
    }
}

</script>

<template>
    <div class="grid grid-cols-5 p-2">
        <input v-model="name" type="text" placeholder="title" class="col-span-1 mr-2 my-0 h-9" />
        <textarea :maxlength="150" v-model="description" type="textarea" @input="resize()" ref="descriptionTextarea"
            placeholder="descripton" class="col-span-1 mr-2 overflow-hidden resize-none h-9 pt-1"></textarea>
        <VueDatePicker v-model="date" range class="max-w-md col-span-1"></VueDatePicker>
        <button v-on:click="addTask($data)" class="justify-self-end col-span-2 text-primary my-auto">add</button>
    </div>
</template>