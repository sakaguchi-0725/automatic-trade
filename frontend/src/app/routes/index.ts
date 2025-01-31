import { HomePage } from "@/pages/home";
import { ReportPage } from "@/pages/report";
import { TradeHistoryPage } from "@/pages/trade-history";
import { createRouter, createWebHistory, Router, RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/trade-history', name: 'trade-history', component: TradeHistoryPage },
  { path: '/report', name: 'report', component: ReportPage }
]

export const createAppRouter = (): Router => {
  return createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes
  })
}