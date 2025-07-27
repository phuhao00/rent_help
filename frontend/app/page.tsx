import Link from 'next/link'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'

export default function HomePage() {
  return (
    <div className="flex flex-col min-h-screen">
      {/* Navigation */}
      <nav className="border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
        <div className="container flex h-14 items-center">
          <div className="mr-4 flex">
            <Link href="/" className="mr-6 flex items-center space-x-2">
              <span className="font-bold text-xl">RentHelp</span>
            </Link>
          </div>
          <div className="flex flex-1 items-center justify-between space-x-2 md:justify-end">
            <nav className="flex items-center space-x-6">
              <Link href="/properties" className="text-sm font-medium">
                房源
              </Link>
              <Link href="/about" className="text-sm font-medium">
                关于我们
              </Link>
            </nav>
            <div className="flex items-center space-x-2">
              <Button variant="ghost" asChild>
                <Link href="/login">登录</Link>
              </Button>
              <Button asChild>
                <Link href="/register">注册</Link>
              </Button>
            </div>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="flex-1 flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100">
        <div className="container px-4 md:px-6">
          <div className="flex flex-col items-center space-y-4 text-center">
            <div className="space-y-2">
              <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl lg:text-6xl">
                找到你的理想住所
              </h1>
              <p className="mx-auto max-w-[700px] text-gray-500 md:text-xl">
                RentHelp 是一个现代化的租房服务平台，连接房东和租客，让租房变得更简单、更透明、更安全。
              </p>
            </div>
            <div className="space-x-4">
              <Button size="lg" asChild>
                <Link href="/properties">浏览房源</Link>
              </Button>
              <Button variant="outline" size="lg" asChild>
                <Link href="/register">发布房源</Link>
              </Button>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-12 md:py-24 lg:py-32">
        <div className="container px-4 md:px-6">
          <div className="flex flex-col items-center justify-center space-y-4 text-center">
            <div className="space-y-2">
              <h2 className="text-3xl font-bold tracking-tighter md:text-4xl">
                为什么选择 RentHelp
              </h2>
              <p className="max-w-[900px] text-gray-500 md:text-xl">
                我们提供安全、便捷、透明的租房服务，让您的租房体验更加愉快。
              </p>
            </div>
          </div>
          <div className="mx-auto grid max-w-5xl items-center gap-6 py-12 lg:grid-cols-3 lg:gap-12">
            <Card>
              <CardHeader>
                <CardTitle>安全可靠</CardTitle>
                <CardDescription>
                  所有房源和用户都经过严格审核，确保信息真实可靠
                </CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-gray-500">
                  我们采用多重验证机制，包括身份认证、房源核实等，为您提供安全的租房环境。
                </p>
              </CardContent>
            </Card>
            <Card>
              <CardHeader>
                <CardTitle>操作便捷</CardTitle>
                <CardDescription>
                  简单直观的界面设计，让您轻松找到心仪的房源
                </CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-gray-500">
                  先进的搜索功能和智能推荐系统，帮您快速找到最适合的房源。
                </p>
              </CardContent>
            </Card>
            <Card>
              <CardHeader>
                <CardTitle>服务全面</CardTitle>
                <CardDescription>
                  从看房到签约，提供全程专业服务支持
                </CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-gray-500">
                  专业的客服团队和完善的售后服务，让您的租房过程无忧无虑。
                </p>
              </CardContent>
            </Card>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t bg-gray-50">
        <div className="container flex flex-col gap-2 sm:flex-row py-6 w-full shrink-0 items-center px-4 md:px-6">
          <p className="text-xs text-gray-500">
            © 2025 RentHelp. 保留所有权利。
          </p>
          <nav className="sm:ml-auto flex gap-4 sm:gap-6">
            <Link href="/terms" className="text-xs hover:underline underline-offset-4">
              服务条款
            </Link>
            <Link href="/privacy" className="text-xs hover:underline underline-offset-4">
              隐私政策
            </Link>
          </nav>
        </div>
      </footer>
    </div>
  )
}
