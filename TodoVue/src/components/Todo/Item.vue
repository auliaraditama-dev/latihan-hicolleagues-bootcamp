<script setup>
import { ref, nextTick } from "vue";

const props = defineProps({
  todo: {
    type: Object,
    required: true,
  },
});
const emit = defineEmits(["ubah", "hapus", "edit"]);
const sedangEdit = ref(false);
const teksEdit = ref("");
const inputEditRef = ref(null);
function mulaiEdit() {
  sedangEdit.value = true;
  teksEdit.value = props.todo.title;
  nextTick(() => {
    inputEditRef.value?.focus();
  });
}
function simpanEdit() {
  if (!sedangEdit.value) return;

  const teksBersih = teksEdit.value.trim();
  if (!teksBersih) {
    batalEdit();
    return;
  }
  emit("edit", {
    id: props.todo.id,
    teksBaru: teksBersih,
  });

  sedangEdit.value = false;
}
function batalEdit() {
  sedangEdit.value = false;
  teksEdit.value = props.todo.title;
}
</script>

<template>
  <li class="todo-item" :class="{ completed: todo.completed }">
    <label class="todo-label" v-if="!sedangEdit">
      <input
        type="checkbox"
        :checked="todo.completed"
        @change="emit('ubah', todo.id)"
      />

      <span class="text" @dblclick="mulaiEdit" title="Klik 2x untuk edit tugas">
        {{ todo.title }}
      </span>
    </label>

    <div v-else class="edit-wrapper">
      <input
        ref="inputEditRef"
        v-model="teksEdit"
        type="text"
        class="input-edit"
        @keydown.enter="simpanEdit"
        @keydown.esc="batalEdit"
        @blur="simpanEdit"
      />
    </div>

    <button
      v-if="!sedangEdit"
      class="btn-delete"
      @click="emit('hapus', todo.id)"
    >
      Hapus
    </button>
  </li>
</template>

<style scoped>
.todo-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 14px;
  background: #faf5ff;
  border: 1px solid #e9d5ff;
  border-radius: 8px;
  margin-bottom: 8px;
  gap: 8px;
}

.todo-label {
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  flex: 1;
}

.text {
  user-select: none;
}

.completed .text {
  text-decoration: line-through;
  color: #a78bfa;
}

.edit-wrapper {
  flex: 1;
}

.input-edit {
  width: 100%;
  padding: 6px 10px;
  border: 1px solid #8b5cf6;
  border-radius: 6px;
  font-size: 14px;
  outline: none;
  box-sizing: border-box;
}

.btn-delete {
  background-color: #ef4444;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 12px;
  cursor: pointer;
}

.btn-delete:hover {
  background-color: #dc2626;
}
</style>
