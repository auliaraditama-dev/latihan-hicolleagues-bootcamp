<script setup>
import { ref, computed, watch, onMounted } from "vue";
import TodoInput from "./components/Todo/Input.vue";
import TodoItem from "./components/Todo/Item.vue";
import TodoFilter from "./components/Todo/Filter.vue";
import TodoStats from "./components/Todo/Stats.vue";

const STORAGE_KEY = "todo_workshop_key";
const todos = ref([]);
onMounted(() => {
  const savedData = localStorage.getItem(STORAGE_KEY);
  if (savedData) {
    try {
      todos.value = JSON.parse(savedData);
    } catch (e) {
      todos.value = [];
    }
  }
});
watch(
  todos,
  (newTodos) => {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(newTodos));
  },
  { deep: true },
);
const filterAktif = ref("semua");

const todosTerfilter = computed(() => {
  if (filterAktif.value === "aktif") {
    return todos.value.filter((todo) => !todo.completed);
  }
  if (filterAktif.value === "selesai") {
    return todos.value.filter((todo) => todo.completed);
  }
  return todos.value;
});

const sisaTugas = computed(() => {
  return todos.value.filter((todo) => !todo.completed).length;
});
function handleTambah(judulBaru) {
  todos.value.push({
    id: Date.now(),
    title: judulBaru,
    completed: false,
  });
}
function handleUbah(id) {
  const item = todos.value.find((todo) => todo.id === id);
  if (item) {
    item.completed = !item.completed;
  }
}
function handleHapus(id) {
  todos.value = todos.value.filter((todo) => todo.id !== id);
}
function handleEdit({ id, teksBaru }) {
  const item = todos.value.find((todo) => todo.id === id);
  if (item) {
    item.title = teksBaru;
  }
}
</script>

<template>
  <main class="app-container">
    <h1>Todo List</h1>

    <TodoInput @tambah="handleTambah" />

    <TodoFilter
      :filter-aktif="filterAktif"
      @gantiFilter="(val) => (filterAktif = val)"
    />

    <ul class="todo-list" v-if="todosTerfilter.length > 0">
      <TodoItem
        v-for="todo in todosTerfilter"
        :key="todo.id"
        :todo="todo"
        @ubah="handleUbah"
        @hapus="handleHapus"
        @edit="handleEdit"
      />
    </ul>
    <p v-else class="empty-text">Tidak ada tugas pada filter ini.</p>

    <TodoStats :sisa-tugas="sisaTugas" :total-tugas="todos.length" />
  </main>
</template>

<style scoped>
:global(body) {
  background-color: #f7f5ff;
}

.app-container {
  max-width: 480px;
  margin: 40px auto;
  padding: 24px;
  background: #ffffff;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(124, 58, 237, 0.1);
  font-family:
    system-ui,
    -apple-system,
    BlinkMacSystemFont,
    "Segoe UI",
    Roboto,
    sans-serif;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
  color: #5b21b6;
}

.todo-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.empty-text {
  text-align: center;
  color: #94a3b8;
  font-size: 14px;
  margin: 24px 0;
}
</style>
