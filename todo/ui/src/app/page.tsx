"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Form, FormField, FormItem, FormMessage } from "@/components/ui/form";
import { FormLabel } from "@/components/ui/form";
import { FormControl } from "@/components/ui/form";
import { yupResolver } from "@hookform/resolvers/yup";
import { Check, Edit, ListIcon, Loader2, Pencil, Plus, Trash2, X } from "lucide-react";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import * as yup from "yup";
import requests from "@/services/api";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Label } from "@/components/ui/label";
import { toast } from "sonner";


const schema = yup.object({
  id: yup.number().nullable().notRequired().default(null),
  title: yup.string().required("Título é obrigatório"),
  completed: yup.boolean().default(false),
  createdAt: yup.date().notRequired().default(null),
  updatedAt: yup.date().notRequired().default(null),
});

type FormData = yup.InferType<typeof schema>;

export default function Home() {
  const [todos, setTodos] = useState<FormData[]>([]);
  const [isLoading, setIsLoading] = useState(false);



  const form = useForm<FormData>({
    resolver: yupResolver(schema),
    defaultValues: {
      title: "",
      completed: false,
    },
  });

  useEffect(() => {
    const fetchTodos = async () => {
      const res = await requests.getTodos();
      console.log(res.data);
      setTodos(res.data);
    };
    fetchTodos();
  }, []);


  const handleDelete = async (id: number) => {
    try {
      setTodos(todos.filter((todo) => todo.id !== id));
      await requests.deleteTodo(id);
      toast.success("Todo excluído com sucesso!", { duration: 10000 });
    } catch (error) {
      console.error(error);
      toast.error("Erro ao excluir todo", { duration: 10000 });
    } finally {
      setIsLoading(false);
    }
  };

  const onSubmitCreate = async (data: FormData) => {
    try {
      setIsLoading(true);
      const res = await requests.createTodo(data)
      setTodos([...todos, res.data]);
      form.reset();
      toast.success("Todo criado com sucesso!", {
        action: {
          label: "Cancelar",
          onClick: () => {
            handleDelete(res.data.id);
          }
        },
        duration: 10000,
      });
    } catch (error) {
      console.error(error);
      toast.error("Erro ao criar todo", { duration: 10000 });
    } finally {
      setIsLoading(false);
    }
  }

  const handleToggleStatus = async (id: number, data: FormData, status: boolean) => {
    try {
      await requests.updateTodo(id, { title: data.title, completed: status });
      setTodos(todos.map((todo) => todo.id === id ? { ...todo, completed: status } : todo));
      toast.success("Status atualizado com sucesso!", { duration: 10000 });
    } catch (error) {
      console.error(error);
      toast.error("Erro ao atualizar status", { duration: 10000 });
    } finally {
      setIsLoading(false);
    }
  }

  const onSubmitUpdate = async (id: number, data: FormData) => {
    try {
      setTodos(todos.map((todo) => todo.id === id ? data : todo));
      await requests.updateTodo(id, { title: data.title, completed: data.completed! });
      form.reset();
      toast.success("Todo atualizado com sucesso!", { duration: 10000 });
    } catch (error) {
      console.error(error);
      toast.error("Erro ao atualizar todo", { duration: 10000 });
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

  const onClickUpdate = (id: number, data: FormData) => {
    form.setValue("id", id);
    form.setValue("title", data.title);
    form.setValue("completed", data.completed);
    form.setValue("createdAt", data.createdAt);
    form.setValue("updatedAt", data.updatedAt);
  }

  return (
    <div className="flex flex-col gap-4 items-center justify-start h-full pt-10">
      <section className="flex items-center justify-center w-full p-4 gap-4">
        <ListIcon className="text-blue-500" size={37} />
        <h3 className="text-2xl font-bold">Tarefas</h3>
      </section>
      <div className="flex w-[750px]  gap-4">
        <section className='flex-1 rounded-md border border-gray-300 p-4'>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="flex gap-2">
              <div className="flex flex-col gap-2">
                <FormField control={form.control} name="title" render={({ field }) => (
                  <FormItem>
                    <FormLabel>Título</FormLabel>
                    <FormControl>
                      <Input type="text" {...field} placeholder="Lavar o carro..." />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )} />
                <FormField
                  control={form.control}
                  name='completed'
                  render={({ field }) => (
                    <FormItem>
                      <FormControl>
                        <RadioGroup onValueChange={(v) => field.onChange(v === "true")} value={field.value ? "true" : "false"}>
                          <div className='flex flex-col gap-2'>
                            <div className='flex items-center space-x-2'>
                              <RadioGroupItem value='false' id='prioridade-instalacao-m' />
                              <Label htmlFor='prioridade-instalacao-m'>Pendente</Label>
                            </div>
                            <div className='flex items-center space-x-2'>
                              <RadioGroupItem value='true' id='prioridade-instalacao-a' />
                              <Label htmlFor='prioridade-instalacao-a'>Concluído</Label>
                            </div>
                          </div>
                        </RadioGroup>
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
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
        <section className='flex-2 rounded-md border border-gray-300 p-4'>
          {
            todos.length === 0 && (
              <p className="text-sm text-gray-500">Nenhum todo encontrado</p>
            )
          }
          <ul className="flex flex-col gap-2">
            {todos.map((todo) => (
              <li key={todo.id} className="flex justify-between items-center cursor-pointer rounded-md">
                <div className="flex items-center gap-2">
                  <span>
                    {
                      todo.completed ? (
                        <Check className="text-green-500" />
                      ) : (
                        <X className="text-red-500" />
                      )
                    }
                  </span>
                  <span className="text-sm overflow-hidden text-ellipsis whitespace-nowrap">{todo.title}</span>
                </div>
                <div className="flex gap-2">
                  <Button variant="outline" onClick={() => handleToggleStatus(todo.id!, todo, !todo.completed)} disabled={isLoading} className="cursor-pointer">
                    {
                      todo.completed ? (
                        <X className="text-red-500" />
                      ) : (
                        <Check className="text-green-500" />
                      )
                    }
                  </Button>
                  <Button variant="outline" onClick={() => onClickUpdate(todo.id!, todo)} disabled={isLoading} className="cursor-pointer">
                    <Edit />
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
