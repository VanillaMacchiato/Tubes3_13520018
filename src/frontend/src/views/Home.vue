<template>
  <div class="home">
    <img alt="logo aplikasi" src="../assets/logo.png" />
    <div class="perkenalan">
      <h1>Selamat datang di Aplikasi DNA At Work!</h1>
    <p id="p0">
      Aplikasi ini terdiri dari beberapa fitur yaitu:
    </p>
    <p id="p1">1. Menerima input penyakit baru</p>
    <p id="p2">2. Memprediksi seseorang menderita penyakit</p>
    <p id="p3">3. Menampilkan urutan hasil prediksi</p>
    <h2> Pilih fitur yang tersedia pada sidebar di sebelah kiri</h2>
    <h2> Semoga aplikasi ini bermanfaat 😃</h2>
    <img id="gambar" alt="gambar bergerak" src="../assets/gambar.gif" />
    <h2> Abaikan gambar di atas </h2>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Home',
  components: {
  }
}
</script>

<style scoped>

#gambar {
  width: 500px;
}
#p0 {
  font-size: 16pt;
}

#p1 {
  margin-right: 70pt;
  font-size: 16pt;
}
#p2 {
  margin-left: 30pt;
  font-size: 16pt;
}
#p3 {
  margin-right: 38pt;
  font-size: 16pt;
}


</style>