<template>
  <div style="padding: 20px; font-family: sans-serif;">
    <section>
      <h3>Latihan 1: Counter</h3>
      <p>Nilai: {{ counter }}</p>
      <button @click="counter++">Tambah</button>
      <button @click="counter--">Kurang</button>
    </section>

    <section style="margin-top: 30px;">
      <h3>Latihan 2: Form Sapaan</h3>
      <input v-model="nama" placeholder="Ketik nama Anda di sini" />
      <p v-if="nama">Halo, selamat datang {{ nama }}!</p>
    </section>

    <section style="margin-top: 30px;">
      <h3>Latihan 3: Kartu Profil</h3>
      <div :style="cardStyle">
        <img :src="avatarUrl" :alt="profilNama" style="width: 100px; border-radius: 50%;" />
        <h4 :title="profilNama">{{ profilNama }}</h4>
        <p :class="pekerjaanClass">{{ pekerjaan }}</p>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const counter = ref(0)

const nama = ref('')

const profilNama = ref('Budi Santoso')
const avatarUrl = ref('https://dummyimage.com/100x100/42b883/ffffff&text=Budi')
const pekerjaan = ref('Software Engineer')
const pekerjaanClass = ref('text-primary')
const cardStyle = ref('border: 1px solid #ddd; padding: 16px; width: 250px; text-align: center; border-radius: 12px; box-shadow: 2px 2px 8px rgba(0,0,0,0.1);')
</script>

<style>
.text-primary {
  font-weight: bold;
  color: #42b883;
}
</style>