<template>
  <nav class="navbar navbar-expand-lg">
    <button
      class="navbar-toggler"
      type="button"
      data-toggle="collapse"
      data-target="#navbarText"
      aria-controls="navbarText"
      aria-expanded="false"
      aria-label="Toggle navigation"
      @click="navopen = !navopen"
    >
      <FontAwesomeIcon
        v-if="!navopen"
        icon="bars"
      />
      <FontAwesomeIcon
        v-if="navopen"
        icon="times"
      />
    </button>

    <div
      id="navbarText"
      class="navbar-collapse"
      :class="{collapse: !navopen}"
    >
      <ul
        v-if="authenticated"
        class="navbar-nav mr-auto"
      >
        <li
          class="nav-item navbar-item"
          @click="navopen = !navopen"
        >
          <router-link
            to="/dashboard"
            class="nav-link"
          >
            {{ $t('dashboard') }}
          </router-link>
        </li>
        <li
          class="nav-item navbar-item"
          @click="navopen = !navopen"
        >
          <router-link
            to="/dashboard/services"
            class="nav-link"
          >
            {{ $t('services') }}
          </router-link>
        </li>
        <li
          v-if="admin"
          class="nav-item navbar-item"
          @click="navopen = !navopen"
        >
          <router-link
            to="/dashboard/users"
            class="nav-link"
          >
            {{ $t('users') }}
          </router-link>
        </li>
        <li
          class="nav-item navbar-item"
          @click="navopen = !navopen"
        >
          <router-link
            to="/dashboard/messages"
            class="nav-link"
          >
            {{ $t('announcements') }}
          </router-link>
        </li>
        <li
          v-if="admin"
          class="nav-item navbar-item"
          @click="navopen = !navopen"
        >
          <router-link
            to="/dashboard/settings"
            class="nav-link"
          >
            {{ $t('settings') }}
          </router-link>
        </li>
        <li
          v-if="admin"
          class="nav-item navbar-item"
          @click="navopen = !navopen"
        >
          <router-link
            to="/dashboard/logs"
            class="nav-link"
          >
            {{ $t('logs') }}
          </router-link>
        </li>
        <li
          v-if="admin"
          class="nav-item navbar-item"
          @click="navopen = !navopen"
        >
          <router-link
            to="/dashboard/help"
            class="nav-link"
          >
            {{ $t('help') }}
          </router-link>
        </li>
      </ul>
      <span class="navbar-text">
        <a
          v-if="authenticated"
          href="#"
          class="nav-link"
          @click.prevent="logout"
        >
          {{ $t('logout') }}
        </a>
        <router-link
          v-else
          to="/login"
          class="nav-link"
        >
          {{ $t('login') }}
        </router-link>
      </span>
    </div>
  </nav>
</template>

<script>
import Api from '../../API';

export default {
    name: 'TopNav',
    props: {
        authenticated: {
            type: Boolean,
            required: true
        }
    },
    data () {
        return {
            navopen: false
        };
    },
    computed: {
        admin () {
            return this.$store.state.admin;
        }
    },
    methods: {
        async logout () {
            await Api.logout();
            this.$store.commit('setHasAllData', false);
            this.$store.commit('setToken', null);
            this.$store.commit('setAdmin', false);
            // this.$cookies.remove("statping_auth")
            await this.$router.push('/logout');
        }
    }
};
</script>
