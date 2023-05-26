import { getUser } from '@src/lib/auth'

export function Profile() {
  const { name, avatarUrl } = getUser()

  return (
    <div className="flex items-center gap-3 text-left transition-colors hover:text-gray-50">
      <img
        alt={`${name} Image`}
        src={avatarUrl}
        className="h-10 w-10 rounded-full"
      />

      <p className="max-w-[200px] text-sm leading-snug">
        {name}
        <a
          href="/api/auth/logout"
          className="block text-red-400 hover:text-red-300"
        >
          Quero sair
        </a>
      </p>
    </div>
  )
}
