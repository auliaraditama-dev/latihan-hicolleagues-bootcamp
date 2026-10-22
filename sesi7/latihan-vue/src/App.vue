<template>
  <div style="padding: 20px; font-family: sans-serif">
    <section>
      <h2>Kartu Biodata</h2>
      <div :style="cardStyle">
        <img
          :src="profilePic"
          :alt="nama"
          style="width: 100px; border-radius: 50%"
        />
        <h3 :title="nama">{{ nama }}</h3>
        <p>
          <strong>Pekerjaan:</strong>
          <span :class="textClass">{{ pekerjaan }}</span>
        </p>
        <p><strong>Umur:</strong> {{ umur }} tahun</p>
      </div>
    </section>

    <hr style="margin: 30px 0" />

    <section>
      <h2>Form Login</h2>
      <form
        @submit.prevent="login"
        style="
          display: flex;
          flex-direction: column;
          max-width: 300px;
          gap: 10px;
        "
      >
        <div>
          <label for="username">Username:</label><br />
          <input
            type="text"
            id="username"
            v-model="username"
            placeholder="Masukkan username"
            style="width: 100%; padding: 5px"
          />
        </div>
        <div>
          <label for="password">Password:</label><br />
          <input
            type="password"
            id="password"
            v-model="password"
            placeholder="Masukkan password"
            style="width: 100%; padding: 5px"
          />
        </div>
        <button type="submit" style="padding: 10px; cursor: pointer">
          Login
        </button>
      </form>

      <div
        v-if="pesanLogin"
        style="
          margin-top: 15px;
          padding: 10px;
          background-color: #eef;
          border-radius: 5px;
        "
      >
        {{ pesanLogin }}
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from "vue";

const nama = ref("Alex Developer");
const pekerjaan = ref("Frontend Engineer");
const umur = ref(25);
const profilePic = ref("https://dummyimage.com/150x150/5c6bc0/ffffff&text=AD");
const cardStyle = ref(
  "border: 1px solid #ddd; padding: 20px; width: 250px; text-align: center; border-radius: 10px; box-shadow: 0 4px 8px rgba(0,0,0,0.1);",
);
const textClass = ref("highlight-text");

const username = ref("");
const password = ref("");
const pesanLogin = ref("");

const login = () => {
  if (username.value && password.value) {
    pesanLogin.value = `Login berhasil! Selamat datang, ${username.value}.`;
  } else {
    pesanLogin.value = "Mohon isi username dan password.";
  }
};
</script>

<style>
.highlight-text {
  color: #5c6bc0;
  font-weight: bold;
}
</style>
