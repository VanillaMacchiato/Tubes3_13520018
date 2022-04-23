<template>
  <div class="home">
    <img alt="logo aplikasi" src="../assets/logo.png" />
    <div class="perkenalan">
      <h1 id="h1">Selamat datang di Aplikasi DNA At Work!</h1>
      <p id="p0">Aplikasi ini terdiri dari beberapa fitur yaitu:</p>
      <p id="p1">1. Menerima input penyakit baru</p>
      <p id="p2">2. Memprediksi seseorang menderita penyakit</p>
      <p id="p3">3. Menampilkan urutan hasil prediksi</p>
      <h2 id="h2">Pilih fitur yang tersedia pada sidebar di sebelah kiri</h2>
      <h2 id="h21">Semoga aplikasi ini bermanfaat 😃</h2>
      <img id="gambar" alt="gambar bergerak" src="../assets/gambar.gif" />
      <h2 id="h22">Abaikan gambar di atas</h2>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Home',
  components: {}
}
</script>

<style scoped>
#gambar {
  width: 500px;
}
#h1 {
  background-color: white;
  width: 630px;
  margin-left: 26pt;
}
#h2 {
  background-color: white;
  width: 590px;
  margin-left: 42pt;
}
#h21 {
  background-color: white;
  width: 400px;
  margin-left: 113pt;
}
#h22 {
  background-color: white;
  width: 270px;
  margin-left: 160pt;
}
#p0 {
  font-size: 16pt;
  background-color: white;
  width: 400px;
  margin-left: 113pt;
}

#p1 {
  margin-left: 112pt;
  font-size: 16pt;
  background-color: white;
  width: 310px;
}
#p2 {
  margin-left: 113pt;
  font-size: 16pt;
  background-color: white;
  width: 442px;
}
#p3 {
  margin-left: 114pt;
  font-size: 16pt;
  background-color: white;
  width: 350px;
}

.perkenalan {
  width: 700px;
  margin: auto;
}

</style>
