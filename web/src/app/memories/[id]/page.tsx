import { api } from '@src/lib/api'
import { ChevronLeft } from 'lucide-react'
import { cookies } from 'next/headers'
import Image from 'next/image'
import Link from 'next/link'

interface Memory {
  id: string
  coverUrl: string
  content: string
  isPublic: boolean
  createAt: string
}

type Props = {
  params: { id: string }
}
export default async function MemoryPage({ params }: Props) {
  const token = cookies().get('token')?.value

  const response = await api.get<Memory>(`/memories/${params.id}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })

  const memory = response.data
  return (
    <div className="flex flex-1 flex-col gap-4 p-8">
      <Link
        href={'/'}
        className="flex items-center gap-1 text-center text-sm text-gray-200 hover:text-gray-100"
      >
        <ChevronLeft className="h-4 w-4" />
        voltar a timeline
      </Link>

      <div className="flex flex-1 flex-col gap-2">
        <Image
          alt=""
          className="aspect-video w-full rounded-lg object-cover"
          width={592}
          height={200}
          src={memory.coverUrl}
        />

        <p className="text-lg leading-relaxed text-gray-100">
          {memory.content}
        </p>
      </div>
    </div>
  )
}
