<script setup>
import { ref, computed } from 'vue'

const tugas = ref([])
const teks = ref('')
const filterSaatIni = ref('Semua')

function tambah() {
  if (!teks.value.trim()) return
  tugas.value.push({ id: Date.now(), teks: teks.value, selesai: false })
  teks.value = ''
}

function hapus(id) {
  tugas.value = tugas.value.filter(t => t.id !== id)
}

function toggleSelesai(t) {
  t.selesai = !t.selesai
}

const daftarTerfilter = computed(() => {
  if (filterSaatIni.value === 'Aktif') {
    return tugas.value.filter(t => !t.selesai)
  } else if (filterSaatIni.value === 'Selesai') {
    return tugas.value.filter(t => t.selesai)
  }
  return tugas.value
})
</script>

<template>
  <div style="max-width: 400px; margin: 20px; font-family: sans-serif;">
    <h3>Daftar Tugas</h3>

    <form @submit.prevent="tambah" style="display: flex; gap: 8px; margin-bottom: 15px;">
      <input v-model="teks" placeholder="Tugas baru..." style="flex: 1; padding: 8px;" />
      <button type="submit" style="padding: 8px 16px;">Tambah</button>
    </form>

    <div style="margin-bottom: 15px; display: flex; gap: 10px;">
      <button @click="filterSaatIni = 'Semua'" :style="{ fontWeight: filterSaatIni === 'Semua' ? 'bold' : 'normal' }">Semua</button>
      <button @click="filterSaatIni = 'Aktif'" :style="{ fontWeight: filterSaatIni === 'Aktif' ? 'bold' : 'normal' }">Aktif</button>
      <button @click="filterSaatIni = 'Selesai'" :style="{ fontWeight: filterSaatIni === 'Selesai' ? 'bold' : 'normal' }">Selesai</button>
    </div>

    <ul style="list-style: none; padding: 0;">
      <li v-for="t in daftarTerfilter" :key="t.id" style="display: flex; justify-content: space-between; margin-bottom: 8px; padding: 8px; border: 1px solid #ccc; align-items: center;">
        <span :style="{ textDecoration: t.selesai ? 'line-through' : 'none', color: t.selesai ? 'gray' : 'black' }">
          {{ t.teks }}
        </span>
        <div>
          <button @click="toggleSelesai(t)" style="margin-right: 5px; cursor: pointer;">
            {{ t.selesai ? 'Batal' : 'Selesai' }}
          </button>
          <button @click="hapus(t.id)" style="color: red; cursor: pointer;">Hapus</button>
        </div>
      </li>
    </ul>
  </div>
</template>