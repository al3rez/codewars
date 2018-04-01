defmodule VowelCount do
  def get_count(str) do
    vowels = ["a", "e", "i", "o", "u"]

    str
    |> String.graphemes()
    |> Enum.filter(fn c -> Enum.member?(vowels, c) end)
    |> Enum.count()
  end
end
