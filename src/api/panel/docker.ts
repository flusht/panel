import { get, post } from '@/utils/request'

export interface ContainerInfo {
    id: string
    names: string[]
    image: string
    state: string
    status: string
    created: number
}

export function getContainerList() {
    return get<ContainerInfo[]>({
        url: '/panel/docker/list',
    })
}

export function startContainer(data: { id: string }) {
    return post<any>({
        url: '/panel/docker/start',
        data,
    })
}

export function stopContainer(data: { id: string }) {
    return post<any>({
        url: '/panel/docker/stop',
        data,
    })
}

export function restartContainer(data: { id: string }) {
    return post<any>({
        url: '/panel/docker/restart',
        data,
    })
}

export function getContainerLogs(id: string) {
    return get<string>({
        url: '/panel/docker/logs',
        data: { id },
    })
}
