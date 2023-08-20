/* eslint-disable @next/next/no-head-element */
import Link from 'next/link';
import './globals.css';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html>
      <body>
        <main>
          <nav>
            <Link href="/">
              Home
            </Link>
            <Link href="/about">
              About
            </Link>
          </nav>
          <div className="page">
            {children}
          </div>
        </main>
      </body>
    </html>
  );
}
