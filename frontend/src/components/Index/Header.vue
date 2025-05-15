<template>
    <div id="header">
        <h1 id="title" class="col-12 text-center pt-4 mt-4 mb-3 header-title font-6">{{core.name}}</h1>
        <h5 id="description" class="col-12 text-center mb-5 header-desc font-3">{{core.description}}</h5>
    </div>
    <div class="mb-3">
      <router-link class="links" :to="admin ? '/dashboard' : '/login'">{{$t('dashboard')}}</router-link>
    </div>
</template>

<script>
export default {
  name: 'Header',
    computed: {
      core() {
          return this.$store.getters.core
      }
    }
}
</script>
