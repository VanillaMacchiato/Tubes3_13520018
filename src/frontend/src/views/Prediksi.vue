<template>
  <div class="prediksi">
    <img alt="text prediksi penyakit" src="../assets/prediksi_penyakit.png" />
    <div class="card">
      <p class="judul">Tes DNA</p>
      <div class="box1">
        <label for="fname" class="label1">Nama Penggguna:</label>
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
          @click="berhasil = false"
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
      </div>
      <div class="namafile">{{ namafile }}</div>
      <button @click="onSubmit" class="submitbtn">
        <i class="fa fa-paper-plane"></i>
        Submit
      </button>
      <p v-if="berhasil">{{ textberhasil }}</p>
      <div class="garis"></div>
      <div class="box4" v-if="!berhasil">
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
      textberhasil: 'Proses Selesai!',
      berhasil: true,
      hasil: ''
    }
  },
  methods: {
    onFileChange(event) {
      var fileData = event.target.files[0]
      this.name = fileData.name
      this.textfile = fileData
    },
    onSubmit() {
      let formData = new FormData()
      formData.append('file', this.textfile)
      formData.append('name', this.namapengguna)
      formData.append('disease', this.namapenyakit)

      fetch('http://localhost:8080/api/v1/predict-patience', {
        method: 'POST',
        body: formData
      })
        .then((res) => {
          return res.json()
        })
        .then((data) => {
          console.log(data)
        })
        .catch((e) => {
          console.log(e)
        })
    }
  }
}
</script>

<style scoped>
.h2 {
  margin-top: 70px;
}

#loading {
  width: 200px;
}

.hasil_tes {
  font-size: 25px;
  color: #15d3fd;
  font-weight: bold;
  margin-top: 60px;
}
.garis {
  background-color: #15d3fd;
  top: 490px;
  height: 2px;
  width: 700px;
  position: absolute;
}
.namafile {
  margin-top: 100px;
  margin-left: 175px;
  position: absolute;
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
  margin-top: 100px;
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
  height: 650px;
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
</style>
