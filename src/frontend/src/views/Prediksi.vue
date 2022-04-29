<template>
  <div class="prediksi">
    <img alt="text prediksi penyakit" src="../assets/prediksi_penyakit.png" />
    <div class="card">
      <p class="judul">Tes DNA</p>
      <div class="box1">
        <label for="fname" class="label1">Nama Pengguna:</label>
        <input
          type="text"
          id="fname"
          v-model="namapengguna"
          name="namapengguna"
          placeholder="Pengguna..."
        />
      </div>
      <div class="box2">
        <p class="sequence">Sequence DNA:</p>
        <label for="file-upload" class="custom-file-upload">
          <i class="fa fa-upload"></i>
          Upload File
        </label>
        <input
          id="file-upload"
          accept=".txt"
          type="file"
          @change="onFileChange"
        />
        <div class="box3">
          <label for="fname" class="label1">Prediksi penyakit:</label>
          <input
            type="text"
            id="fname"
            v-model="namapenyakit"
            name="namapenyakit"
            placeholder="Penyakit..."
          />
        </div>
        <p class="textpilihan">Pilih algoritma pencarian:</p>
        <div class="algo1">
          <label class="container">
            Boyer-Moore
            <input
              type="radio"
              checked="checked"
              name="radio"
              value="bm"
              v-model="radiobtn"
            />
            <span class="btnRadio"></span>
          </label>
        </div>
        <div class="algo2">
          <label class="container">
            Knuth-Morris-Pratt
            <input type="radio" name="radio" value="kmp" v-model="radiobtn" />
            <span class="btnRadio"></span>
          </label>
        </div>
      </div>
      <div class="namafile">{{ namafile }}</div>
      <button @click="onSubmit" class="submitbtn">
        <i class="fa fa-paper-plane"></i>
        Submit
      </button>
      <p>{{ textberhasil }}</p>
      <div class="garis"></div>
      <div class="box4" v-if="!selesai">
        <img id="loading" alt="img loading" src="../assets/loading_fast.gif" />
        <p class="judul">Sedang Diproses</p>
      </div>
      <p class="hasil_tes" v-if="berhasil">Hasil Tes</p>
      <h2 class="h2" v-if="berhasil">{{ hasil }}</h2>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Prediksi',
  components: {},
  data() {
    return {
      namafile: '',
      namapenyakit: '',
      namapengguna: '',
      textfile: '',
      textberhasil: '',
      berhasil: false,
      selesai: true,
      hasil: '',
      radiobtn: ''
    }
  },
  methods: {
    onFileChange(event) {
      var fileData = event.target.files[0]
      this.namafile = fileData.name
      this.textfile = fileData
      this.berhasil = false
      this.textberhasil = ''
    },
    onSubmit() {
      this.textberhasil = ''
      this.selesai = false
      let formData = new FormData()
      formData.append('file', this.textfile)
      formData.append('name', this.namapengguna)
      formData.append('disease', this.namapenyakit)
      if (this.radiobtn == 'bm') {
        formData.append('algorithm', 'BoyerMoore')
      } else {
        formData.append('algorithm', 'KMP')
      }
      fetch('https://dna-at-work-backend.herokuapp.com/api/v1/predict-patience', {
        method: 'POST',
        body: formData
      })
        .then((res) => {
          return res.json()
        })
        .then((data) => {
          console.log(data)
          this.hasil =
            data.data.patientName +
            ' - ' +
            data.data.diseaseName +
            ' - ' +
            data.data.hasDisease.toString().toUpperCase() +
            ' - ' +
            data.data.likeness +
            '%'
          this.berhasil = true
          this.textberhasil = 'Proses Selesai!'
          this.selesai = true
        })
        .catch((e) => {
          console.log(e)
          this.textberhasil = 'Proses Gagal!'
          this.selesai = true
        })
    }
  }
}
</script>

<style scoped>
.h2 {
  margin-top: 70px;
}
.textpilihan {
  margin-top: 30px;
  font-size: 15px;
  font-weight: bold;
}
#loading {
  width: 200px;
}
.pilihan {
  margin-top: 30px;
}
.algo1 {
  margin-left: 255px;
  width: 140px;
  height: 40px;
}
.algo2 {
  margin-left: 255px;
  width: 185px;
}
.hasil_tes {
  font-size: 25px;
  color: #15d3fd;
  font-weight: bold;
  margin-top: 60px;
}
.garis {
  background-color: #15d3fd;
  top: 550px;
  height: 2px;
  width: 700px;
  position: absolute;
}
.namafile {
  margin-top: 10px;
  position: absolute;
  margin-left: 300px;
}
.judul {
  font-size: 25px;
  color: #15d3fd;
  font-weight: bold;
}
.label1 {
  margin-right: 50px;
  font-size: 15px;
  font-weight: bold;
}
input[type='text'],
select {
  width: 180px;
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
.submitbtn {
  background-color: DodgerBlue;
  border: none;
  color: white;
  padding: 12px 30px;
  cursor: pointer;
  font-size: 20px;
  left: 960px;
  top: 480px;
  margin-top: 40px;
}
.submitbtn:hover {
  background-color: RoyalBlue;
}

.box1 {
  width: 300px;
  height: 50px;
  position: absolute;
}
.box3 {
  width: 200px;
  height: 50px;
  top: 225px;
  position: absolute;
  margin-left: 500px;
}
.box4 {
  margin-top: 50px;
}
.card {
  background-color: #fefefe;
  box-shadow: inset 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  transition: 0.3s;
  padding: 16px;
  margin: auto;
  width: 700px;
  height: 730px;
  border-radius: 14px;
  border: 1px solid #15d3fd;
}
.card:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
}

input[type='file'] {
  display: none;
}
.custom-file-upload {
  border: 1px solid #ccc;
  display: inline-block;
  padding: 6px 12px;
  cursor: pointer;
  left: 15px;
}
.sequence {
  font-size: 15px;
  font-weight: bold;
  margin-right: 15px;
}
.box2 {
  margin-left: 10px;
}

.container {
  display: block;
  position: relative;
  padding-left: 35px;
  margin-bottom: 12px;
  cursor: pointer;
  font-size: 18px;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.container input {
  position: absolute;
  opacity: 0;
  cursor: pointer;
}

.btnRadio {
  position: absolute;
  top: 0;
  left: 0;
  height: 25px;
  width: 25px;
  background-color: #eee;
  border-radius: 50%;
}

.container:hover input ~ .btnRadio {
  background-color: #ccc;
}

.container input:checked ~ .btnRadio {
  background-color: #2196f3;
}

.btnRadio:after {
  content: '';
  position: absolute;
  display: none;
}

.container input:checked ~ .btnRadio:after {
  display: block;
}

.container .btnRadio:after {
  top: 9px;
  left: 9px;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: white;
}
</style>
