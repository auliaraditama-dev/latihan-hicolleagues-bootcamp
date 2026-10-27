<script setup>
defineProps({
  filterAktif: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['gantiFilter'])

const opsiFilter = [
  { label: 'Semua', value: 'semua' },
  { label: 'Aktif', value: 'aktif' },
  { label: 'Selesai', value: 'selesai' }
]
</script>

<template>
  <div class="filter-group">
    <button
      v-for="opsi in opsiFilter"
      :key="opsi.value"
      class="btn-filter"
      :class="{ aktif: filterAktif === opsi.value }"
      @click="emit('gantiFilter', opsi.value)"
    >
      {{ opsi.label }}
    </button>
  </div>
</template>

<style scoped>
.filter-group {
  display: flex;
  justify-content: center;
  gap: 8px;
  margin-bottom: 20px;
}

.btn-filter {
  padding: 6px 14px;
  background-color: #f5f3ff;
  border: 1px solid #ddd6fe;
  border-radius: 6px;
  font-size: 13px;
  font-weight: 500;
  color: #5b21b6;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-filter:hover {
  background-color: #ede9fe;
}

.btn-filter.aktif {
  background-color: #7c3aed;
  border-color: #7c3aed;
  color: #ffffff;
}
</style>