<template>
  <div class="p-4">
    <n-card title="Docker Container Management">
      <template #header-extra>
        <n-button size="small" @click="refreshList" :loading="loading">
          Refresh
        </n-button>
      </template>

      <n-data-table
        :columns="columns"
        :data="containerList"
        :loading="loading"
        :row-key="(row) => row.id"
      />
    </n-card>

    <n-modal v-model:show="showLogs" title="Container Logs" preset="card" style="width: 800px; height: 600px;">
      <n-card title="Logs" :bordered="false" size="small">
        <n-log :log="logsContent" :rows="25" />
      </n-card>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { h, onMounted, ref } from 'vue'
import { NCard, NDataTable, NButton, NTag, NSpace, NModal, NLog, useMessage } from 'naive-ui'
import { getContainerList, startContainer, stopContainer, restartContainer, getContainerLogs, type ContainerInfo } from '@/api/panel/docker'

const message = useMessage()
const loading = ref(false)
const containerList = ref<ContainerInfo[]>([])
const showLogs = ref(false)
const logsContent = ref('')

const columns = [
  {
    title: 'Name',
    key: 'names',
    render(row: ContainerInfo) {
      return row.names.map(n => n.replace(/^\//, '')).join(', ')
    }
  },
  {
    title: 'Image',
    key: 'image'
  },
  {
    title: 'State',
    key: 'state',
    render(row: ContainerInfo) {
      return h(
        NTag,
        {
          type: row.state === 'running' ? 'success' : 'error',
          bordered: false
        },
        { default: () => row.state }
      )
    }
  },
  {
    title: 'Status',
    key: 'status'
  },
  {
    title: 'Actions',
    key: 'actions',
    render(row: ContainerInfo) {
      return h(NSpace, null, {
        default: () => [
          h(
            NButton,
            {
              size: 'tiny',
              type: 'success',
              disabled: row.state === 'running',
              onClick: () => handleStart(row.id)
            },
            { default: () => 'Start' }
          ),
          h(
            NButton,
            {
              size: 'tiny',
              type: 'warning',
              disabled: row.state !== 'running',
              onClick: () => handleStop(row.id)
            },
            { default: () => 'Stop' }
          ),
          h(
            NButton,
            {
              size: 'tiny',
              type: 'info',
              onClick: () => handleRestart(row.id)
            },
            { default: () => 'Restart' }
          ),
          h(
            NButton,
            {
              size: 'tiny',
              onClick: () => handleLogs(row.id)
            },
            { default: () => 'Logs' }
          )
        ]
      })
    }
  }
]

const refreshList = async () => {
  loading.value = true
  try {
    const res = await getContainerList()
    if (res.code === 0) {
      containerList.value = res.data
    } else {
      message.error(res.msg || 'Failed to get container list')
    }
  } catch (error) {
    message.error('Network error')
  } finally {
    loading.value = false
  }
}

const handleStart = async (id: string) => {
  try {
    const res = await startContainer({ id })
    if (res.code === 0) {
      message.success('Container started')
      refreshList()
    } else {
      message.error(res.msg || 'Failed to start')
    }
  } catch (error) {
    message.error('Error starting container')
  }
}

const handleStop = async (id: string) => {
  try {
    const res = await stopContainer({ id })
    if (res.code === 0) {
      message.success('Container stopped')
      refreshList()
    } else {
      message.error(res.msg || 'Failed to stop')
    }
  } catch (error) {
    message.error('Error stopping container')
  }
}

const handleRestart = async (id: string) => {
  try {
    const res = await restartContainer({ id })
    if (res.code === 0) {
      message.success('Container restarted')
      refreshList()
    } else {
      message.error(res.msg || 'Failed to restart')
    }
  } catch (error) {
    message.error('Error restarting container')
  }
}

const handleLogs = async (id: string) => {
  try {
    const res = await getContainerLogs(id)
    if (res.code === 0) {
      logsContent.value = res.data
      showLogs.value = true
    } else {
      message.error(res.msg || 'Failed to get logs')
    }
  } catch (error) {
    message.error('Error getting logs')
  }
}

onMounted(() => {
  refreshList()
})
</script>
