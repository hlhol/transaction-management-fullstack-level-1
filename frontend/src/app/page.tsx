'use client'

import { useState, useEffect } from 'react'

type Transaction = {
  transaction_id: string
  account_id: string
  amount: number
  balance: number 
}

export default function Home() {
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [accountId, setAccountId] = useState('')
  const [amount, setAmount] = useState('')

  const fetchTransactions = async () => {
    const res = await fetch('/transactions')
    const data = await res.json()
    setTransactions(Array.isArray(data) ? data.reverse() : [])
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!accountId || !amount) return

    const amountNum = parseInt(amount)
    if (isNaN(amountNum)) return

    const res = await fetch('/transactions', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ account_id: accountId, amount: amountNum }),
    })

    if (!res.ok) {
      alert('Failed to submit transaction')
      return
    }

    // Refresh transactions
    await fetchTransactions()

    setAccountId('')
    setAmount('')
  }

  useEffect(() => {
    fetchTransactions()
  }, [])

  return (
    <main className="max-w-md mx-auto p-6 font-sans">
      <h1 className="text-2xl font-bold mb-6">Accounting App</h1>

      {/* Form */}
      <form onSubmit={handleSubmit} className="space-y-4 mb-8">
        <input
          data-type="account-id"
          type="text"
          placeholder="Account ID"
          value={accountId}
          onChange={e => setAccountId(e.target.value)}
          className="w-full p-2 border rounded"
          required
        />
        <input
          data-type="amount"
          type="number"
          placeholder="Amount"
          value={amount}
          onChange={e => setAmount(e.target.value)}
          className="w-full p-2 border rounded"
          required
        />
        <input
          data-type="transaction-submit"
          type="submit"
          value="Submit"
          className="w-full bg-blue-600 text-white py-2 rounded cursor-pointer"
        />
      </form>

      {/*  list */}
      <div>
        {transactions.map((tx) => (
          <div
            key={tx.transaction_id}
            data-type="transaction"
            data-account-id={tx.account_id}
            data-amount={tx.amount}
            data-balance={tx.balance} 
            className="border p-4 rounded mb-4 bg-gray-50"
          >
            <p><strong>Account ID:</strong> {tx.account_id}</p>
            <p><strong>Amount:</strong> {tx.amount}</p>
            <p><strong>Balance after transaction:</strong> {tx.balance}</p>
          </div>
        ))}
      </div>
    </main>
  )
}