<template>
  <div class="tentang">
    <img alt="text mengenai aplikasi" src="../assets/mengenai_aplikasi.png" />
    <h1>Halaman ini berisi informasi sangat penting dalam aplikasi ini</h1>
    <div class="kotak">
      <div class="card">
        <p id="p0">Aplikasi ini dibuat oleh:</p>
        <p id="p1">1. Bariza Haqi 13520018</p>
        <p id="p2">2. Farrel Farandieka Fibriyanto 13520054</p>
        <p id="p3">3. Vito Ghifari 13520153</p>
      </div>
    </div>
    <div class="akhir">
      <h2>Sudah sih gitu aja</h2>
      <h2>Tehe</h2>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Tentang',
  components: {}
}
</script>

<style scoped>
#p0 {
  font-size: 16pt;
}

#p1 {
  margin-left: 15pt;
  font-size: 16pt;
}
#p2 {
  margin-left: 20pt;
  font-size: 16pt;
}
#p3 {
  margin-left: 12pt;
  font-size: 16pt;
}
.list-nama {
  margin-top: 120pt;
}
.akhir {
  margin-top: 100pt;
}
.card {
  box-shadow: inset 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  transition: 0.3s;
  padding: 16px;
  margin: auto;
  width: 500px;
  height: 200px;
  border-radius: 14px;
  border: 1px solid #15d3fd;
  background-color: #fefefe;
}
.card:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
}

.kotak {
  margin-top: 100pt;
}
</style>
