<script>
import SidebarLink from './SidebarLink'
import { collapsed, toggleSidebar, sidebarWidth } from './state'

export default {
  props: {},
  components: { SidebarLink },
  setup() {
    return { collapsed, toggleSidebar, sidebarWidth }
  }
}
</script>

<template>
  <div class="sidebar" :style="{ width: sidebarWidth }">
    <h1>
      <span v-if="collapsed">
        <div>D</div>
        <div>M</div>
      </span>
      <span v-else>DNA Matching</span>
    </h1>

    <SidebarLink to="/" icon="fas fa-home">Halaman Utama</SidebarLink>
    <SidebarLink to="/input" icon="fas fa-heart">Input Penyakit</SidebarLink>
    <SidebarLink to="/prediksi" icon="fas fa-thermometer-quarter">Prediksi Penyakit</SidebarLink>
    <SidebarLink to="/hasil" icon="fas fa-user-circle">Hasil Prediksi</SidebarLink>
    <SidebarLink to="/tentang" icon="fas fa-user">Mengenai Aplikasi</SidebarLink>

    <span
      class="collapse-icon"
      :class="{ 'rotate-180': collapsed }"
      @click="toggleSidebar"
    >
      <i class="fas fa-angle-double-left" />
    </span>
  </div>
</template>

<style>
:root {
  --sidebar-bg-color: #13c0f0;
  --sidebar-item-hover: #119cd8;
  --sidebar-item-active: #0aaaf4;
}
</style>

<style scoped>
.sidebar {
  color: white;
  background-color: var(--sidebar-bg-color);

  float: left;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  bottom: 0;
  padding: 0.5em;

  transition: 0.3s ease;

  display: flex;
  flex-direction: column;
}

.sidebar h1 {
  height: 2.5em;
}

.collapse-icon {
  position: absolute;
  bottom: 0;
  padding: 0.75em;

  color: rgba(255, 255, 255, 0.7);

  transition: 0.2s linear;
}

.rotate-180 {
  transform: rotate(180deg);
  transition: 0.2s linear;
}
</style>
