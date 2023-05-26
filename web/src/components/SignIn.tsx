import { User } from 'lucide-react'
import Link from 'next/link'

export function SignIn() {
  return (
    <Link
      href={`https://github.com/login/oauth/authorize?client_id=${process.env.NEXT_PUBLIC_GITHUB_CILENT_ID}`}
      className="flex items-center gap-3 text-left transition-colors hover:text-gray-50"
    >
      <div className="grid h-10 w-10 place-content-center rounded-full bg-gray-400">
        <User className="h-5 w-5 text-gray-500" />
      </div>

      <p className="max-w-[200px] text-sm leading-snug">
        <span className="underline">Crie sua conta </span> e salve suas
        memorias!
      </p>
    </Link>
  )
}
