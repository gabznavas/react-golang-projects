"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Form, FormField, FormItem, FormMessage } from "@/components/ui/form";
import { FormLabel } from "@/components/ui/form";
import { FormControl } from "@/components/ui/form";
import { yupResolver } from "@hookform/resolvers/yup";
import { Loader2, Pencil, Plus, Trash2 } from "lucide-react";
import { useContext, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import * as yup from "yup";
import requests from "@/services/api";


const schema = yup.object({
  id: yup.number().nullable().notRequired().default(null),
  title: yup.string().required("Título é obrigatório"),
  completed: yup.boolean().notRequired().default(false),
  createdAt: yup.date().notRequired().default(null),
  updatedAt: yup.date().notRequired().default(null),
});

type FormData = yup.InferType<typeof schema>;

export default function Home() {
  const [todos, setTodos] = useState<FormData[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  const form = useForm<FormData>({
    resolver: yupResolver(schema),
  });

  useEffect(() => {
    const fetchTodos = async () => {
      const res = await requests.getTodos();
      console.log(res.data);
      setTodos(res.data);
    };
    fetchTodos();
  }, []);

  const onSubmitCreate = async (data: FormData) => {
    try {
      setIsLoading(true);
      const res = await requests.createTodo(data)
      setTodos([...todos, res.data]);
      form.reset();
    } catch (error) {
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  }

  const onSubmitUpdate = async (id: number, data: FormData) => {
    try {
      setTodos(todos.map((todo) => todo.id === id ? data : todo));
      await requests.updateTodo(id, data);
      form.reset();
    } catch (error) {
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  }

  const onSubmit = async (data: FormData) => {
    const id = form.getValues("id");
    if (id) {
      onSubmitUpdate(id, data);
    } else {
      onSubmitCreate(data);
    }
  }

  const handleDelete = async (id: number) => {
    try {
      setTodos(todos.filter((todo) => todo.id !== id));
      await requests.deleteTodo(id);
    } catch (error) {
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };


  const onClickUpdate = (id: number, data: FormData) => {
    form.setValue("id", id);
    form.setValue("title", data.title);
    form.setValue("completed", data.completed);
    form.setValue("createdAt", data.createdAt);
    form.setValue("updatedAt", data.updatedAt);
  }

  return (
    <div className="flex flex-col gap-4 items-center justify-start  h-full">
      <h1 className="text-2xl font-bold">Todos</h1>
      <div className="flex gap-4">
        <section className='rounded-md border border-gray-300 p-4 w-72'>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="flex gap-2">
              <div className="flex flex-col gap-2">
                <FormField control={form.control} name="title" render={({ field }) => (
                  <FormItem>
                    <FormLabel>Título</FormLabel>
                    <FormControl>
                      <Input type="text" {...field} placeholder="Add a new todo" />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )} />
                <Button type="submit" disabled={isLoading} className="cursor-pointer">
                  {
                    isLoading ? (
                      <Loader2 className="animate-spin" />
                    ) : (
                      <>
                        {
                          form.getValues("id") ? (
                            <Pencil />
                          ) : (
                            <Plus />
                          )
                        }
                        {form.getValues("id") ? "Atualizar" : "Adicionar"}
                      </>
                    )
                  }
                </Button>
                {
                  form.watch("id") && (
                    <Button variant="outline" type="button" onClick={() => form.reset({
                      id: null,
                      title: "",
                      completed: false,
                      createdAt: null,
                      updatedAt: null,
                    })} disabled={isLoading}>
                      <Trash2 />
                      Limpar
                    </Button>
                  )
                }
              </div>
            </form>
          </Form>
        </section>
        <section className='rounded-md border border-gray-300 p-4 w-80'>
          {
            todos.length === 0 && (
              <p className="text-sm text-gray-500">Nenhum todo encontrado</p>
            )
          }
          <ul className="flex flex-col gap-2">
            {todos.map((todo) => (
              <li key={todo.id} className="flex justify-between items-center">
                <span className="text-sm overflow-hidden text-ellipsis whitespace-nowrap">{todo.title}</span>
                <div className="flex gap-2">
                  <Button variant="default" onClick={() => onClickUpdate(todo.id!, todo)} disabled={isLoading} className="cursor-pointer">
                    <Pencil />
                  </Button>
                  <Button variant="destructive" onClick={() => handleDelete(todo.id!)} disabled={isLoading} className="cursor-pointer">
                    <Trash2 />
                  </Button>
                </div>
              </li>
            ))}
          </ul>
        </section>
      </div>
    </div>
  );
}
