import { defineStore } from 'pinia'
import type { RouteRecordRaw } from 'vue-router'
import type { SiteConfig } from '/@/stores/interface'

export const useSiteConfig = defineStore('siteConfig', {
    state: (): SiteConfig => {
        return {
            siteName: '4HFire',
            recordNumber: '渝ICP备2020013067号-2',
            version: 'v1.0.0',
            cdnUrl: 'https://demo.buildadmin.com',
            apiUrl: '',
            upload: {
                maxsize: 10485760,
                savename: "\/storage\/{topic}\/{year}{mon}{day}\/{filename}{filesha1}{.suffix}",
                mimetype: "jpg,png,bmp,jpeg,gif,webp,zip,rar,xls,xlsx,doc,docx,wav,mp4,mp3,txt",
                mode: "local"
            },
            headNav: [],
            initialize: false,
            userInitialize: false,
        }
    },
    actions: {
        dataFill(state: SiteConfig) {
            // this.$state = state
        },
        setHeadNav(headNav: RouteRecordRaw[]) {
            this.headNav = headNav
        },
        setInitialize(initialize: boolean) {
            this.initialize = initialize
        },
        setUserInitialize(userInitialize: boolean) {
            this.userInitialize = userInitialize
        },
    },
})
