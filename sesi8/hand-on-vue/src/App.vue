<script setup>
import { ref, computed, watch, onMounted } from 'vue'

const tugas = ref([])
const teks = ref('')
const cari = ref('')
const pesanTersimpan = ref(false)

onMounted(() => {
  const simpanan = localStorage.getItem('dataTugas')
  if (simpanan) {
    tugas.value = JSON.parse(simpanan)
  }
})

watch(tugas, (baru) => {
  localStorage.setItem('dataTugas', JSON.stringify(baru))
  
  pesanTersimpan.value = true
  setTimeout(() => {
    pesanTersimpan.value = false
  }, 1500)
}, { deep: true })

function tambah() {
  if (!teks.value.trim()) return
  tugas.value.push({ id: Date.now(), teks: teks.value })
  teks.value = ''
}

function hapus(id) {
  tugas.value = tugas.value.filter(t => t.id !== id)
}

const hasil = computed(() => {
  return tugas.value.filter(t =>
    t.teks.toLowerCase().includes(cari.value.toLowerCase())
  )
})

const jumlahTugas = computed(() => tugas.value.length)
</script>

<template>
  <div style="max-width: 400px; margin: 20px; font-family: sans-serif;">
    <h3>Daftar Tugas</h3>
    
    <div style="margin-bottom: 15px;">
      <input v-model="cari" placeholder="Cari tugas..." style="width: 100%; padding: 8px;" />
    </div>

    <form @submit.prevent="tambah" style="display: flex; gap: 8px; margin-bottom: 15px;">
      <input v-model="teks" placeholder="Tugas baru..." style="flex: 1; padding: 8px;" />
      <button type="submit" style="padding: 8px 16px;">Tambah</button>
    </form>

    <div v-show="pesanTersimpan" style="color: green; font-size: 12px; margin-bottom: 10px;">
      Tersimpan!
    </div>

    <ul style="list-style: none; padding: 0;">
      <li v-for="t in hasil" :key="t.id" style="display: flex; justify-content: space-between; margin-bottom: 8px; padding: 8px; border: 1px solid #ccc;">
        {{ t.teks }}
        <button @click="hapus(t.id)" style="color: red; cursor: pointer;">Hapus</button>
      </li>
    </ul>

    <p v-if="hasil.length === 0" style="color: gray; font-style: italic;">
      Tidak ada hasil.
    </p>

    <p style="margin-top: 15px; font-size: 14px; font-weight: bold;">
      Total tugas: {{ jumlahTugas }}
    </p>
  </div>
</template>