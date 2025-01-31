<script lang="ts" setup>
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <aside class="fixed pt-6 w-60 h-screen bg-sidebar">
    <div class="h-full px-3 px-4 overflow-y-auto">
      <div class="text-white text-3xl mb-6 p-2">Logo</div>
      <ui class="space-y-2 font-medium text-white font-normal">
        <li class="list-none">
          <RouterLink :to="{ name: 'home' }"
            :class="{ 'selected-list font-bold': route.name === 'home' }"
            class="flex items-center p-2 rounded-lg hover:font-bold group"
          >
            <span class="ms-3">ホーム</span>
          </RouterLink>
        </li>
        <li class="list-none">
          <RouterLink :to="{ name: 'trade-history' }"
            :class="{ 'selected-list font-bold': route.name === 'trade-history' }"
            class="flex items-center p-2 rounded-lg hover:font-bold group"
          >
            <span class="ms-3">取引履歴</span>
          </RouterLink>
        </li>
        <li class="list-none">
          <RouterLink :to="{ name: 'report' }"
            :class="{ 'selected-list font-bold': route.name === 'report' }"
            class="flex items-center p-2 rounded-lg hover:font-bold group"
          >
            <span class="ms-3">レポート</span>
          </RouterLink>
        </li>
      </ui>
    </div>
  </aside>
</template>

<style lang="css" scoped>
.bg-sidebar {
  background-color: #2B2D34;
}
.selected-list {
  background-color: #15AF8A;
}
</style>