import { EmptyMemories } from '@src/components/EmptyMemories'
import { api } from '@src/lib/api'
import { getUser } from '@src/lib/auth'
import dayjs from 'dayjs'
import ptBr from 'dayjs/locale/pt-br'
import { ArrowRight } from 'lucide-react'
import { Metadata } from 'next'
import { cookies } from 'next/headers'
import Image from 'next/image'
import Link from 'next/link'

dayjs.locale(ptBr)

interface Memory {
  id: string
  coverUrl: string
  excerpt: string
  createAt: string
}

export function generateMetadata(): Metadata {
  const isAuthenticated = cookies().has('token')
  if (isAuthenticated) {
    const { name } = getUser()
    return {
      title: `${name}  |  NLW Spacetime`,
    }
  }
  return {}
}
export default async function Home() {
  const isAuthenticated = cookies().has('token')
  if (!isAuthenticated) {
    return <EmptyMemories />
  }

  const token = cookies().get('token')?.value

  const response = await api.get<Memory[]>('/memories', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })

  const memories = response.data
  if (memories.length === 0) {
    return <EmptyMemories />
  }
  return (
    <div className="flex flex-col gap-10 p-8">
      {memories.map((memory) => (
        <div key={memory.id} className="space-y-4">
          <time className="-ml-8 flex items-center gap-2 text-sm text-gray-100 before:h-px before:w-5 before:bg-gray-50">
            {dayjs(memory.createAt).format('D[ de ]MMMM[, ]YYYY')}
          </time>
          <Image
            alt=""
            className="aspect-video w-full rounded-lg object-cover"
            width={592}
            height={200}
            src={memory.coverUrl}
          />

          <p className="text-lg leading-relaxed text-gray-100">
            {memory.excerpt}
          </p>
          <Link
            href={`/memories/${memory.id}`}
            className="flex items-center gap-2 text-sm text-gray-200 hover:text-gray-100"
          >
            Ler mais <ArrowRight className="h-4 w-4" />
          </Link>
        </div>
      ))}
    </div>
  )
}
