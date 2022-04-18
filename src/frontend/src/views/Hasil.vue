<template>
  <div class="hasil">
    <img alt="text hasil prediksi" src="../assets/hasil_prediksi.png" />
  <div class="card">
    <p class="judul"> Cari Hasil Prediksi</p>
    <div class="box1">
      <label for="fname" class="label1"></label>
    <input type="text" id="fname" v-model="nama" name="namapenyakit" placeholder="<tanggal_prediksi><spasi><nama_penyakit>">
    </div>
    <button @click="berhasil = true" class="searchbtn">
          <i class="fa fa-search"></i> Cari
        </button>
    <p v-if="berhasil"> {{textberhasil}} </p>
  </div>
  <div class="box2" v-if="!berhasil">
    <div id="img">
    <img alt="loading" src="../assets/loading.gif" />
    </div>
    <p class="judul"> Sedang Melakukan Pencarian... </p>
  </div>
  <div class="card1" v-if="berhasil">
    <p class="judul"> {{nama}} </p>
  </div>
  </div>
</template>

<script>

export default {
  name: 'Hasil',
  components: {
  },
  data() {
    return {
      nama: "",
      textberhasil: "Berhasil ditemukan!" ,
      berhasil: false,
    };
  },
  methods: {
  },
  
  }

</script>

<style scoped>
.judul {
  font-size: 25px;
  color: #15d3fd;
  font-weight: bold;
}
input[type=text], select {
  width: 310px;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

.buttons {
  text-align:justify;
}
.searchbtn {
  background-color: DodgerBlue;
  border: none;
  color: white;
  padding: 12px 30px;
  cursor: pointer;
  font-size: 20px;
  left: 960px;
  top: 480px;
  margin-top: 80px;

}
.searchbtn:hover {
  background-color: RoyalBlue;
}

.box1 {
  width: 200px;
  height: 50px;
  position: absolute;
  margin-left:100px;
 
}

.card {
  background-color: #fefefe;
  box-shadow: inset 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  transition: 0.3s;
  padding: 16px;
  margin :auto;
  width: 500px;
  height: 250px;
  border-radius: 14px;
  border: 1px solid #15d3fd;
}
.card:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
}

.card1 {
  background-color: #fefefe;
  box-shadow: inset 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  transition: 0.3s;
  padding: 16px;
  margin :auto;
  width: 500px;
  height: 400px;
  border-radius: 14px;
  border: 1px solid #15d3fd;
  margin-top: 100px;
}
.card1:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
}

</style>