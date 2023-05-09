<template>
    <div>
        创建一个简单的页面，这个页面。
        注意：要在端系统页面上 提前在路由中配置好，才能访问到。并且在页面上给角色设置好权限。
    </div>
</template>

<script>
// 导入testApi
import {testApi} from '@/api/testApi.js'
export default {
    name: "whatever",
    created() {
        // 调用testApi接口。 注意: 先增加api权限；再给角色设置权限。增加api时，最前面的/不要忘记。
        testApi()
    },
}
</script>