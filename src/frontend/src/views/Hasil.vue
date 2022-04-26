<template>
  <div class="hasil">
    <img alt="text hasil prediksi" src="../assets/hasil_prediksi.png" />
    <div class="card">
      <p class="judul">Cari Hasil Prediksi</p>
      <div class="box1">
        <label for="fname" class="label1"></label>
        <input
          type="text"
          id="fname"
          v-model="inputUser"
          name="namapenyakit"
          placeholder="<tanggal_prediksi><spasi><nama_penyakit>"
        />
      </div>
      <button @click="onSubmit" class="searchbtn">
        <i class="fa fa-search"></i>
        Cari
      </button>
      <p v-if="berhasil">{{ textberhasil }}</p>
    </div>
    <div class="box2">
      <div class="mencari" v-if="!selesai">
        <div id="img">
          <img alt="loading" src="../assets/loading.gif" />
        </div>
        <p class="judul">Sedang Melakukan Pencarian...</p>
      </div>
    </div>
    <div class="card1" v-if="berhasil">
      <p class="judul">{{ nama }}</p>
      <div class="container" v-for="user in hasil" :key="user.id">
        <div class="box3">
          {{ user.tanggal }} - {{ user.nama }} - {{ user.penyakit }} -
          {{ user.terkena }} - {{ user.kemiripan }}%
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Hasil',
  components: {},
  data() {
    return {
      inputUser: '',
      nama: '',
      textberhasil: '',
      berhasil: false,
      selesai: true,
      hasil: []
    }
  },
  methods: {
    onSubmit() {
      this.berhasil = false
      this.selesai = false
      this.hasil = []

      this.url =
        'https://dna-at-work.netlify.app/api/v1/result' +
        '/?input=' +
        this.inputUser.replaceAll(' ', '%20')
      fetch(this.url)
        .then((res) => {
          return res.json()
        })
        .then((data) => {
          console.log(data)
          for (let i = 0; i < data.data.length; i++) {
            let entry = {
              tanggal: data.data[i].Date.slice(0, 10),
              nama: data.data[i].PatientName,
              penyakit: data.data[i].DiseaseName,
              terkena: data.data[i].HasDisease.toString().toUpperCase(),
              kemiripan: data.data[i].Likeness
            }
            this.hasil.push(entry)
          }
          this.berhasil = true
          this.textberhasil = 'Berhasil ditemukan!'
        })
        .catch((e) => {
          console.log(e)
          this.textberhasil = 'Hasil tidak ditemukan!'
        })
      this.selesai = true
    }
  }
}
</script>

<style scoped>
.box2 {
  margin: auto;
  width: 500px;
}
.box3 {
  padding: 16px;
  margin: auto;
  width: 400px;
  height: 20px;
  border-radius: 14px;
  border: 1px solid #15d3fd;
  margin-top: 20px;
}
.judul {
  font-size: 25px;
  color: #15d3fd;
  font-weight: bold;
  background-color: white;
  width: 375px;
  margin-left: 47pt;
}
input[type='text'],
select {
  width: 310px;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

.buttons {
  text-align: justify;
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
  margin-left: 100px;
}

.card {
  background-color: #fefefe;
  box-shadow: inset 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  transition: 0.3s;
  padding: 16px;
  margin: auto;
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
  margin: auto;
  width: 500px;
  height: 100%;
  border-radius: 14px;
  border: 1px solid #15d3fd;
  margin-top: 100px;
}
.card1:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
}
</style>
