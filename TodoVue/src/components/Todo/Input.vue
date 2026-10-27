<script setup>
import { ref } from "vue";

const emit = defineEmits(["tambah"]);
const judul = ref("");
const pesanError = ref("");

function handleSubmit() {
  const teksBersih = judul.value.trim();
  if (!teksBersih) {
    pesanError.value = "Tugas tidak boleh kosong, ya!";
    return;
  }
  emit("tambah", teksBersih);
  judul.value = "";
  pesanError.value = "";
}
function handleInput() {
  if (pesanError.value) {
    pesanError.value = "";
  }
}
</script>

<template>
  <div class="input-wrapper">
    <form @submit.prevent="handleSubmit" class="todo-input-form">
      <input
        v-model="judul"
        type="text"
        placeholder="Tulis tugas baru..."
        class="input-box"
        :class="{ 'input-error': pesanError }"
        @input="handleInput"
      />
      <button type="submit" class="btn-add">Tambah</button>
    </form>

    <p v-if="pesanError" class="error-msg">⚠️ {{ pesanError }}</p>
  </div>
</template>

<style scoped>
.input-wrapper {
  margin-bottom: 16px;
}

.todo-input-form {
  display: flex;
  gap: 8px;
}

.input-box {
  flex: 1;
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 14px;
  outline: none;
  transition:
    border-color 0.2s ease,
    box-shadow 0.2s ease;
}

.input-box:focus {
  border-color: #7c3aed;
}

.input-box.input-error {
  border-color: #ef4444;
  background-color: #fef2f2;
}

.btn-add {
  padding: 10px 18px;
  background-color: #7c3aed;
  color: white;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
}

.btn-add:hover {
  background-color: #6d28d9;
}

.error-msg {
  color: #dc2626;
  font-size: 12px;
  margin: 6px 0 0 2px;
  font-weight: 500;
  animation: fadeIn 0.2s ease-in-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-4px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
