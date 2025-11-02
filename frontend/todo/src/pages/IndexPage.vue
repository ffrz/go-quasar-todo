<template>
  <q-page class="row items-center justify-evenly">
    <div>
      <div class="text-subtitle2">TODOS:</div>
      <div v-for="todo in todos" :key="todo.id">
        <div
          class="todo row items-center"
          :class="todo.done ? 'completed' : 'incomplete'"
          @click="toggleTodo(todo)"
        >
          <q-icon name="crop_square" class="q-mr-sm" />
          <div class="label">
            {{ todo.content }}
          </div>
          <q-icon v-if="todo.done" name="check" class="q-ml-sm" />
        </div>
      </div>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { Todo } from 'components/models';

const todos = ref<Todo[]>([
  {
    id: 1,
    content: 'Learn Go',
    done: false,
  },
  {
    id: 2,
    content: 'Check Emails',
    done: false,
  },
  {
    id: 3,
    content: 'Learn Advance Quasar',
    done: false,
  },
  {
    id: 4,
    content: 'Advance Scalable Architecture Exploration',
    done: false,
  },
  {
    id: 5,
    content: 'Eat well',
    done: false,
  },
]);

const toggleTodo = (todo: Todo) => {
  todo.done = !todo.done;
};
</script>
<style scoped>
.todo {
  cursor: pointer;
}
.completed .label {
  color: #666;
  text-decoration: line-through;
}
</style>
