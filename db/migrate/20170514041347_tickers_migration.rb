class TickersMigration < ActiveRecord::Migration[5.1]
  def change
    create_table :tickers do |t|
      t.integer :last, limit: 4, null: false, unsigned: true
      t.integer :bid, limit: 4, null: false, unsigned: true
      t.integer :ask, limit: 4, null: false, unsigned: true
      t.integer :high, limit: 4, null: false, unsigned: true
      t.integer :low, limit: 4, null: false, unsigned: true
      t.float :volume, null: false
      t.timestamps null: true
    end
  end
end
