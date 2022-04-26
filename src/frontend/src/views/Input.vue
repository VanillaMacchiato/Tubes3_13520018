<template>
  <div class="input">
    <img alt="text input penyakit" src="../assets/input_penyakit.png" />
    <div class="card">
      <p class="judul">Tambahkan Penyakit</p>
      <div class="box1">
        <label for="fname" class="label1">Nama penyakit:</label>
        <input
          type="text"
          id="fname"
          v-model="nama"
          name="namapenyakit"
          placeholder="Penyakit..."
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
      </div>
      <div class="namefile">{{ namafile }}</div>
      <button @click="onSubmit" class="submitbtn">
        <i class="fa fa-paper-plane"></i>
        Submit
      </button>
      <p>{{ textberhasil }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Input',
  components: {},
  data() {
    return {
      namafile: '',
      nama: '',
      textfile: '',
      textberhasil: ''
    }
  },
  methods: {
    onFileChange(event) {
      var fileData = event.target.files[0]
      this.namafile = fileData.name
      this.textfile = fileData
      this.textberhasil = ''
    },
    onSubmit() {
      let formData = new FormData()
      formData.append('disease-name', this.nama)
      formData.append('file', this.textfile)

      fetch('https://dna-at-work-backend.herokuapp.com/api/v1/add-disease', {
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
      //if berhasil
      this.textberhasil = 'Berhasil ditambahkan!'
      //else
      this.textberhasil =
        'Gagal ditambahkan! Data sudah ada atau input nama kosong'
    }
  }
}
</script>

<style scoped>
.sequence {
  font-size: 15px;
  font-weight: bold;
  margin-right: 15px;
}
.namefile {
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
  margin-right: 70px;
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
  margin-top: 120px;
}
.submitbtn:hover {
  background-color: RoyalBlue;
}

.box1 {
  width: 200px;
  height: 50px;
  position: absolute;
}

.card {
  background-color: #fefefe;
  box-shadow: inset 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  transition: 0.3s;
  padding: 16px;
  margin: auto;
  width: 500px;
  height: 350px;
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
.box2 {
  margin-left: 330px;
}
</style>
