<script lang="ts" setup>
import { SideMenu } from '@/shared/ui';
import { RouterView, RouterLink, useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <SideMenu />
  <div class="pt-6 pb-4 px-4 sm:ml-64">
    <RouterView />
  </div>
</template>