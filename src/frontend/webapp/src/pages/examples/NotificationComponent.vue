<script setup>
// imports
import { onMounted, ref } from "vue";
import { api } from "boot/axios";
import { useQuasar } from "quasar";
// variables
const $q = useQuasar();
const data = ref(null);
const alerts = [
  {
    color: "negative",
    message: "Woah! Danger! You are getting good at this!",
    icon: "bi-bug-fill",
  },
  {
    message: "You need to know about this!",
    icon: "bi-exclamation-diamond-fill",
  },
  { message: "Wow! Nice job!", icon: "bi-hand-thumbs-up" },
  { color: "teal", message: "Quasar is cool! Right?", icon: "bi-messenger" },
  {
    color: "purple",
    message: "Jim just pinged you",
    avatar: "https://cdn.quasar.dev/img/boy-avatar.png",
  },
  {
    multiLine: true,
    message:
      "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Hic quisquam non ad sit assumenda consequuntur esse inventore officia. Corrupti reiciendis impedit vel, fugit odit quisquam quae porro exercitationem eveniet quasi.",
  },
];

function loadData() {
  api
    .get("/api/backend")
    .then((response) => {
      data.value = response.data;
    })
    .catch(() => {
      $q.notify({
        color: "negative",
        position: "top",
        message: "Loading failed",
        icon: "bi-bug-fill",
      });
    });
}

// functions
onMounted(() => {
  loadData();
});
function showNotif(position) {
  const { color, textColor, multiLine, icon, message, avatar } =
    alerts[Math.floor(Math.random() * 10) % alerts.length];
  const random = Math.random() * 100;

  const twoActions = random > 70;
  const buttonColor = color ? "white" : void 0;

  $q.notify({
    color,
    textColor,
    icon: random > 30 ? icon : null,
    message,
    position,
    avatar,
    multiLine,
    actions: twoActions
      ? [
        {
          label: "Reply",
          color: buttonColor,
          handler: () => {
            /* console.log('wooow') */
          },
        },
        {
          label: "Dismiss",
          color: "yellow",
          handler: () => {
            /* console.log('wooow') */
          },
        },
      ]
      : random > 40
        ? [
          {
            label: "Reply",
            color: buttonColor,
            handler: () => {
              /* console.log('wooow') */
            },
          },
        ]
        : null,
    timeout: Math.random() * 5000 + 3000,
  });
}
//
function showNotifyBtn() {
  $q.notify({
    message: "Jim pinged you.",
    caption: "5 minutes ago",
    color: "secondary",
  });
}
</script>

<template>
  <div class="fixed fixed-center">
  <h4>Notification</h4>
  <div class="q-pa-md q-gutter-y-sm column items-center">

    <div>
      <div class="row q-gutter-sm">
        <q-btn round size="sm" color="secondary" @click="showNotif('top-left')">
          <q-icon name="bi-arrow-bar-left" class="rotate-45" />
        </q-btn>
        <q-btn round size="sm" color="accent" @click="showNotif('top')">
          <q-icon name="bi-arrow-bar-up" />
        </q-btn>
        <q-btn
          round
          size="sm"
          color="secondary"
          @click="showNotif('top-right')"
        >
          <q-icon name="bi-arrow-bar-up" class="rotate-45" />
        </q-btn>
      </div>
    </div>

    <div>
      <div class="row q-gutter-sm">
        <div>
          <q-btn round size="sm" color="accent" @click="showNotif('left')">
            <q-icon name="bi-arrow-bar-left" />
          </q-btn>
        </div>
        <div>
          <q-btn round size="sm" color="accent" @click="showNotif('center')">
            <q-icon name="bi-arrow-repeat" />
          </q-btn>
        </div>
        <div>
          <q-btn round size="sm" color="accent" @click="showNotif('right')">
            <q-icon name="bi-arrow-bar-right" />
          </q-btn>
        </div>
      </div>
    </div>

    <div>
      <div class="row q-gutter-sm">
        <div>
          <q-btn
            round
            size="sm"
            color="secondary"
            @click="showNotif('bottom-left')"
          >
            <q-icon name="bi-arrow-bar-right" class="rotate-135" />
          </q-btn>
        </div>
        <div>
          <q-btn round size="sm" color="accent" @click="showNotif('bottom')">
            <q-icon name="bi-arrow-bar-down" />
          </q-btn>
        </div>
        <div>
          <q-btn
            round
            size="sm"
            color="secondary"
            @click="showNotif('bottom-right')"
          >
            <q-icon name="bi-arrow-bar-right" class="rotate-45" />
          </q-btn>
        </div>
      </div>
    </div>
  </div>
  <div class="q-pa-md q-gutter-y-sm column items-center">
    <q-btn color="purple" @click="showNotifyBtn" label="Show with caption" />
  </div>
  <h4>Badge</h4>
  <div class="q-pa-md q-gutter-md">
    <q-badge color="blue">
      #4D96F2
    </q-badge>

    <q-badge color="orange" text-color="black" label="2" />

    <q-badge color="purple">
      <q-icon name="bi-bluetooth" color="white" />
    </q-badge>

    <q-badge color="red">
      12 <q-icon name="bi-exclamation-diamond" color="white" class="q-ml-xs" />
    </q-badge>

    <q-badge color="primary">v1.0.0+</q-badge>

    <div>
      Feature <q-badge color="primary">v1.0.0+</q-badge>
    </div>

    <q-item
      clickable
      v-ripple
      class="rounded-borders"
      :class="$q.dark.isActive ? 'bg-grey-9 text-white' : 'bg-grey-2'"
    >
      <q-item-section avatar>
        <q-avatar rounded>
          <img src="https://cdn.quasar.dev/img/chaosmonkey.png" alt="chaosmonkey">
        </q-avatar>
      </q-item-section>

      <q-item-section>
        <q-item-label>
          Ganglia
        </q-item-label>
        <q-item-label caption>
          <q-badge color="yellow-6" text-color="black">
            3
            <q-icon
              name="bi-exclamation-diamond"
              size="14px"
              class="q-ml-xs"
            />
          </q-badge>
        </q-item-label>
      </q-item-section>

      <q-item-section side>
        <span>2 min ago</span>
      </q-item-section>
    </q-item>
    <q-badge rounded color="yellow" />
    <q-badge rounded color="green" />
    <q-badge rounded color="red" />
    <div class="q-gutter-md q-ml-none">
      <q-btn round icon="bi-bell">
        <q-badge floating color="red" rounded />
      </q-btn>
      <q-btn color="blue">
        Notifications
        <q-badge color="red" rounded floating />
      </q-btn>
    </div>
    <div>
      <q-badge color="blue" rounded class="q-mr-sm" />Status
    </div>
  </div>
  <h4>Progress-Bar (AjaxBAr)</h4>
  <div>
    <q-btn color="primary" label="Trigger" @click="trigger" />
  </div>
    <h4>Circular Progress</h4>
    <div class="q-pa-md flex flex-center">
      <q-circular-progress
        indeterminate
        rounded
        size="50px"
        color="lime"
        class="q-ma-md"
      />

      <q-circular-progress
        indeterminate
        size="90px"
        :thickness="0.2"
        color="lime"
        center-color="grey-8"
        track-color="transparent"
        class="q-ma-md"
      />

      <q-circular-progress
        indeterminate
        size="45px"
        :thickness="1"
        color="grey-8"
        track-color="lime"
        class="q-ma-md"
      />

      <q-circular-progress
        indeterminate
        size="50px"
        :thickness="0.22"
        rounded
        color="lime"
        track-color="grey-3"
        class="q-ma-md"
      />

      <q-circular-progress
        indeterminate
        size="75px"
        :thickness="0.6"
        color="lime"
        center-color="grey-8"
        class="q-ma-md"
      />

      <q-circular-progress
        indeterminate
        size="40px"
        :thickness="0.4"
        font-size="50px"
        color="lime"
        track-color="grey-3"
        center-color="grey-8"
        class="q-ma-md"
      />
    </div>
  </div>

</template>

<style scoped></style>
