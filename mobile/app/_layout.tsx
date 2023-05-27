import { ImageBackground } from 'react-native'
import { SplashScreen, Stack } from 'expo-router'

import {
  useFonts,
  Roboto_400Regular,
  Roboto_700Bold,
} from '@expo-google-fonts/roboto'
import { BaiJamjuree_700Bold } from '@expo-google-fonts/bai-jamjuree'
import { styled } from 'nativewind'
import { StatusBar } from 'expo-status-bar'
import { useEffect, useState } from 'react'
import * as SecureStore from 'expo-secure-store'

import blurBg from '../src/assets/blur-bg.png'
import Stripes from '../src/assets/stripes.svg'

const StyledStipes = styled(Stripes)

export default function Layout() {
  const [isUserAuthenticated, setIsUserAuthenticated] = useState<
    null | boolean
  >(null)
  const [hasLoadedFonts] = useFonts({
    Roboto_400Regular,
    Roboto_700Bold,
    BaiJamjuree_700Bold,
  })

  useEffect(() => {
    async function loadData() {
      const token = await SecureStore.getItemAsync('token')
      setIsUserAuthenticated(!!token)
    }

    loadData()
  }, [])

  if (!hasLoadedFonts || isUserAuthenticated === null) {
    return <SplashScreen />
  }

  return (
    <ImageBackground
      source={blurBg}
      className="relative flex-1 bg-gray-900"
      imageStyle={{
        position: 'absolute',
        left: '-100%',
      }}
    >
      <StyledStipes className="absolute left-2" />
      <StatusBar style="light" translucent />
      <Stack
        screenOptions={{
          headerShown: false,
          animation: 'fade',
          contentStyle: { backgroundColor: 'transparent' },
        }}
      >
        <Stack.Screen name="index" redirect={isUserAuthenticated} />
        <Stack.Screen name="memories/index" />
        <Stack.Screen name="memories/new" />
        <Stack.Screen name="memories/[id]" />
      </Stack>
    </ImageBackground>
  )
}
