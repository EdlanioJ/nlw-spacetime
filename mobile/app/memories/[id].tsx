/* eslint-disable jsx-a11y/alt-text */
import { Link, SplashScreen, useSearchParams } from 'expo-router'
import { useEffect, useState } from 'react'
import { ScrollView, Text, TouchableOpacity, View } from 'react-native'
import { Image } from 'expo-image'
import { useSafeAreaInsets } from 'react-native-safe-area-context'
import dayjs from 'dayjs'
import ptBR from 'dayjs/locale/pt-br'
import * as SecureStore from 'expo-secure-store'
import Icon from '@expo/vector-icons/Feather'

import NLWLogo from '../../src/assets/nlw-logo.svg'
import { api } from '../../src/lib/api'

dayjs.locale(ptBR)

interface Memory {
  id: string
  coverUrl: string
  content: string
  isPublic: boolean
  createAt: string
}

export default function MemoryDetails() {
  const [memory, setMemory] = useState<Memory | null>(null)
  const { bottom, top } = useSafeAreaInsets()
  const { id } = useSearchParams()

  useEffect(() => {
    async function loadMemory() {
      const token = await SecureStore.getItemAsync('token')
      const response = await api.get<Memory>(`memories/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      setMemory(response.data)
    }

    loadMemory()
  }, [id])

  if (memory === null) {
    return <SplashScreen />
  }

  return (
    <ScrollView
      className="flex-1 px-8"
      contentContainerStyle={{
        paddingBottom: bottom,
        paddingTop: top,
        gap: 24,
      }}
    >
      <View className="mt-4 flex-row items-center justify-between">
        <NLWLogo />

        <Link href={'/memories/index'} asChild>
          <TouchableOpacity className="h-10 w-10 items-center justify-center rounded-full bg-purple-500">
            <Icon name="arrow-left" size={16} color={'#fff'} />
          </TouchableOpacity>
        </Link>
      </View>
      <View className="flex-row items-center gap-2">
        <View className="h-px w-5 bg-gray-50" />
        <Text className="font-body text-xs text-gray-100">
          {dayjs(memory.createAt).format('D[ de ]MMMM[, ]YYYY')}
        </Text>
      </View>
      <Image
        source={memory.coverUrl}
        className="h-32 w-full rounded-lg object-cover"
      />
      <Text className="font-body text-base leading-relaxed text-gray-100">
        {memory.content}
      </Text>
    </ScrollView>
  )
}
